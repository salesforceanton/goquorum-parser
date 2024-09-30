package transport

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/subjects"
)

type Basis interface {
	Client() *Client
	Logger() *slog.Logger
	Timeout() time.Duration
	Exit() <-chan struct{}
}

type RequesterFunc[Req any, Resp any] func(b Basis, ctx context.Context, req Req) (Resp, error)

// Responder instance function.
// If async is true, use semaphore inside processor to control workload.
// Call of function starts goroutine.
type ResponderFunc[Req any, Resp any] func(b Basis, async bool,
	processor func(ctx context.Context, req Req) (Resp, error),
) error

type EmitterFunc[Event any] func(b Basis, ctx context.Context, event Event) error

// Receiver instance function.
// If async is true, use semaphore inside processor to control workload.
// Call of function starts goroutine.
type ReceiverFunc[Event any] func(b Basis, async bool, processor func(ctx context.Context, ev Event)) error

// NewRequester construct function to make request and get response from some subject.
func NewRequester[Req any, Resp any, ReqPtr interface {
	*Req
	ProtocolMessage
}, RespPtr interface {
	*Resp
	ProtocolMessage
}](
	subject subjects.Subject,
	tagName string,
) RequesterFunc[Req, Resp] {
	return func(b Basis, ctx context.Context, req Req) (Resp, error) {
		client := b.Client()

		var empty Resp

		reqPtr := ReqPtr(&req)

		reqBz, err := Marshal(reqPtr)
		if err != nil {
			b.Logger().ErrorContext(ctx, "Failed to encode request", "err", err, "tag", tagName)
			return empty, err
		}

		timeoutCtx, cancel := context.WithTimeout(ctx, b.Timeout())
		defer cancel()

		msg := nats.Msg{
			Subject: subject,
			Data:    reqBz,
		}

		if client.nc.HeadersSupported() {
			propagateTraceID(timeoutCtx, &msg)
		}

		natsResp, err := client.nc.RequestMsgWithContext(timeoutCtx, &msg)
		if err != nil {
			b.Logger().ErrorContext(ctx, "Failed during request", "err", err, "tag", tagName)
			return empty, err
		}

		resp := RespPtr(new(Resp))

		err = ParseResponse(natsResp, resp)
		if err != nil {
			b.Logger().ErrorContext(ctx, "Error in response", "err", err, "tag", tagName)
			return empty, err
		}

		return *resp, nil
	}
}

func NewResponder[Req any, Resp any, ReqPtr interface {
	*Req
	ProtocolMessage
}, RespPtr interface {
	*Resp
	ProtocolMessage
}](
	subject subjects.Subject,
	tagName string,
) ResponderFunc[Req, Resp] {
	// return instance of function
	return func(b Basis, async bool,
		processor func(ctx context.Context, req Req) (Resp, error),
	) error {
		// declare responder func
		responder := func(b Basis, msg *nats.Msg) {
			var err error

			// unmarshal
			req := ReqPtr(new(Req))

			if len(msg.Data) > 0 {
				err = Unmarshal(msg.Data, req)
				if err != nil {
					b.Logger().Error("Failed to unmarshal request", "err", err, "tag", tagName)
					return
				}
			}

			ctx := restoreTraceID(context.Background(), msg)

			// call processor for request
			// reply only if reply subject is not empty
			resp, err := processor(ctx, *req)
			if err != nil {
				if msg.Reply == "" {
					return
				}
				b.Client().RespondWithError(msg, err)
				return
			}

			if msg.Reply == "" {
				return
			}

			// marshal response
			respPtr := RespPtr(&resp)
			respData, err := Marshal(respPtr)
			if err != nil {
				if msg.Reply == "" {
					return
				}
				b.Client().RespondWithError(msg, err)
				return
			}

			err = msg.Respond(respData)
			if err != nil {
				b.Logger().Error("Failed to respond", "err", err, "subject", subject)
			}
		}

		client := b.Client()
		sub, err := client.nc.SubscribeSync(subject)
		if err != nil {
			return err
		}

		go func() {
			for {
				select {
				case <-b.Exit():
					err := sub.Unsubscribe()
					if err != nil {
						b.Logger().Error("Failed to unsubscribe subject", "err", err, "subject", subject)
					}
					return
				default:
					msg, err := sub.NextMsg(b.Timeout())
					if err != nil {
						if errors.Is(err, nats.ErrTimeout) && sub.IsValid() {
							continue
						}

						b.Logger().Error("Failed to get next message. Exit", "err", err, "subject", subject)
						return
					}

					if async {
						go responder(b, msg)
					} else {
						responder(b, msg)
					}
				}
			}
		}()

		return nil
	}
}

// NewReceiver construct new event receiver.
func NewReceiver[Event any, EventPtr interface {
	*Event
	ProtocolMessage
}](
	subject subjects.Subject,
	tagName string,
) ReceiverFunc[Event] {
	return func(b Basis, async bool, processor func(ctx context.Context, ev Event)) error {
		client := b.Client()
		sub, err := client.nc.SubscribeSync(subject)
		if err != nil {
			return err
		}

		go func() {
			for {
				select {
				case <-b.Exit():
					err := sub.Unsubscribe()
					if err != nil {
						b.Logger().Error("Failed to unsubscribe subject", "err", err, "subject", subject)
					}
					return
				default:
					msg, err := sub.NextMsg(b.Timeout())
					if err != nil {
						if errors.Is(err, nats.ErrTimeout) && sub.IsValid() {
							continue
						}

						b.Logger().Error("Failed to get next message. Exit", "err", err, "subject", subject)
						return
					}

					event := EventPtr(new(Event))

					err = Unmarshal(msg.Data, event)
					if err != nil {
						b.Logger().Error("Failed to unmarshal event", "err", err, "subject", subject)
						continue
					}

					ctx := context.Background()

					if async {
						go processor(ctx, *event)
					} else {
						processor(ctx, *event)
					}
				}
			}
		}()

		return nil
	}
}

func NewEmitter[Event any, EventPtr interface {
	*Event
	ProtocolMessage
}](
	subject subjects.Subject,
	tagName string,
) EmitterFunc[Event] {
	return func(b Basis, ctx context.Context, event Event) error {
		client := b.Client()

		eventPtr := EventPtr(&event)

		eventBz, err := Marshal(eventPtr)
		if err != nil {
			b.Logger().Error("Failed to encode event", "err", err, "subject", subject)
			return err
		}

		msg := nats.Msg{
			Subject: subject,
			Data:    eventBz,
		}

		return client.nc.PublishMsg(&msg)
	}
}
