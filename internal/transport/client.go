package transport

import (
	"log/slog"
	"net"
	"time"

	"github.com/nats-io/nats.go"
)

// Client nats client.
type Client struct {
	nc     *nats.Conn
	logger *slog.Logger

	config   Config // saved config need for reconnection
	flagStop bool   // flag for graceful shutdown
}

// Config is NATS Client config struct.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Stream   string // client turns on JetStream if only Stream != ""
}

func New(config Config, logger *slog.Logger) (*Client, error) {
	c := Client{
		config: config,
		logger: logger,
	}

	err := c.connect()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Client) connect() error {
	var err error

	// regular NATS client
	c.nc, err = nats.Connect(
		net.JoinHostPort(c.config.Host, c.config.Port),
		nats.UserInfo(c.config.User, c.config.Password),
		nats.RetryOnFailedConnect(true),
		nats.ReconnectWait(time.Second),
	)

	return err
}

// Stop stops nats client.
func (c *Client) Stop() error {
	c.flagStop = true
	return nil
}

func (c *Client) Publish(subject string, message []byte) error {
	err := c.nc.Publish(subject, message)
	if err != nil {
		return err
	}

	return nil
}
