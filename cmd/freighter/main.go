package main

import (
	"context"
	"os"

	"freighter.dev/go/freighter/cmd/freighter/cli"
	"freighter.dev/go/freighter/internal/flags"
	"freighter.dev/go/freighter/pkg/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.NewLogger(os.Stdout)
	ctx = logger.WithContext(ctx)

	if err := cli.New(ctx, &flags.CliRootOpts{}).ExecuteContext(ctx); err != nil {
		logger.Errorf("%v", err)
		cancel()
		os.Exit(1)
	}
}
