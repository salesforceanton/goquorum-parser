package transport

import (
	"errors"
	"strconv"

	"github.com/nats-io/nats.go"

	custerr "github.com/salesforceanton/goquorum-parser/internal/custerror"
)

const (
	ErrorHeader = "err"
)

var (
	ErrWrongHeaderFormat = errors.New("wrong error header format")
)

func ErrorFromMsg(msg *nats.Msg) error {
	errHeader := msg.Header.Values(ErrorHeader)

	if len(errHeader) == 0 {
		return nil
	}

	if len(errHeader) != 4 {
		return ErrWrongHeaderFormat
	}

	httpCode, err := strconv.ParseInt(errHeader[0], 10, 64)
	if err != nil {
		return err
	}

	customCode, err := strconv.ParseUint(errHeader[3], 10, 16)
	if err != nil {
		return err
	}

	codeExpr := errHeader[1]
	message := errHeader[2]

	switch httpCode {
	case custerr.BadRequestErrorCode:
		return custerr.NewBadRequestError(uint16(customCode), codeExpr, message)
	case custerr.UnauthorizedErrorCode:
		return custerr.NewUnauthorizedError(uint16(customCode), codeExpr, message)
	case custerr.ForbiddenErrorCode:
		return custerr.NewForbiddenError(uint16(customCode), codeExpr, message)
	case custerr.NotFoundErrorCode:
		return custerr.NewNotFoundError(uint16(customCode), codeExpr, message)
	case custerr.DataConflictErrorCode:
		return custerr.NewDataConflictError(uint16(customCode), codeExpr, message)
	case custerr.ValidationErrorCode:
		return custerr.NewValidationError(uint16(customCode), codeExpr, message)
	case custerr.ServerErrorCode:
		return custerr.NewServerError(errors.New(message))
	default:
		return custerr.NewGeneralError(uint16(httpCode), codeExpr, message)
	}
}

// ParseResponse parse response from nats request.
func ParseResponse(msg *nats.Msg, value ProtocolMessage) error {
	err := ErrorFromMsg(msg)
	if err != nil {
		return err
	}

	err = Unmarshal(msg.Data, value)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RespondWithError(msg *nats.Msg, err error) {
	var httpCode, customCode uint16
	var errMsg, codeExpr string

	custErr, ok := err.(custerr.CustomError)
	if ok {
		httpCode = custErr.HTTP()
		codeExpr = custErr.CodeExpr()
		errMsg = custErr.Error()
		customCode = custErr.CustomCode()
	} else {
		httpCode = custerr.ServerErrorCode
		errMsg = "internal server error: " + err.Error()
	}

	msg.Header = nats.Header{
		ErrorHeader: []string{
			strconv.Itoa(int(httpCode)),
			codeExpr,
			errMsg,
			strconv.Itoa(int(customCode)),
		},
	}

	err = msg.RespondMsg(msg)
	if err != nil {
		c.logger.Error("Failed send response to transport layer", "err", err.Error())
		return
	}
}
