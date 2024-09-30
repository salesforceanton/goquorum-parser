package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/salesforceanton/goquorum-parser/services/backend"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	module, err := backend.New()
	if err != nil {
		panic(err)
	}

	err = module.Start(ctx)
	if err != nil {
		panic(err)
	}

	<-ctx.Done()

	err = module.Stop()
	if err != nil {
		panic(err)
	}
}
