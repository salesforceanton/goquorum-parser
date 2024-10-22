package main

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	httpTimeout       = time.Minute
	rpcRequestTimeout = time.Second * 7
)

type Restarter struct {
	rpcUrl        string
	period        time.Duration
	httpClient    *ethclient.Client
	logger        *slog.Logger
	currentHeight uint64

	exit chan struct{}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// TODO: read from config file
	const rpcUrl = "http://localhost:8545"
	const period = time.Minute

	restarter, err := NewRestarter(rpcUrl, period)
	if err != nil {
		panic(err)
	}

	go restarter.Start()

	<-ctx.Done()

	restarter.Stop()
}

func NewRestarter(rpcUrl string, period time.Duration) (*Restarter, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{ //nolint:exhaustruct
		AddSource: true,
		Level:     slog.LevelDebug,
	})).With("service", "restarter")

	restater := &Restarter{
		rpcUrl: rpcUrl,
		period: period,
		logger: logger,
		exit:   make(chan struct{}),
	}

	err := restater.initClient()
	if err != nil {
		return nil, err
	}

	return restater, nil
}

func (r *Restarter) Start() {
	r.logger.Debug("start service")

	ticker := time.NewTicker(r.period)
	defer ticker.Stop()

	for {
		select {
		case <-r.exit:
			return
		case <-ticker.C:
			bcHeight, err := r.httpClient.BlockNumber(context.Background())
			if err != nil {
				r.logger.Error("Failed to get current height", "err", err)
				continue
			}

			// case if network is stopped
			if bcHeight == r.currentHeight {
				r.logger.Debug("restart network...")
				r.restartNetwork()
			} else {
				r.currentHeight = bcHeight
				r.logger.Debug(fmt.Sprintf("block height: %d", r.currentHeight))
			}
		}
	}
}

func (r *Restarter) Stop() {
	r.logger.Debug("stop service")
	close(r.exit)
}

func (r *Restarter) initClient() error {
	rpcClient, err := rpc.DialHTTPWithClient(
		r.rpcUrl,
		&http.Client{
			Timeout: rpcRequestTimeout,
		},
	)
	if err != nil {
		r.logger.Error(fmt.Sprintf("error with http connect %s", r.rpcUrl))
		return err
	}

	r.httpClient = ethclient.NewClient(rpcClient)
	r.logger.Debug(fmt.Sprintf("successfully connected by http to %s", r.rpcUrl))
	return nil
}

func (r *Restarter) restartNetwork() error {
	cmd := exec.Command("bash", "./restart.sh")
	cmd.Stdin = strings.NewReader("")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	r.logger.Debug(fmt.Sprintf("Output: %s", out.String()))
	return nil
}
