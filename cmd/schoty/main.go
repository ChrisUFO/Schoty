package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ChrisUFO/Schoty/internal/logging"
	"github.com/ChrisUFO/Schoty/internal/ui"
	"github.com/charmbracelet/bubbletea"
)

var (
	helpFlag     = flag.Bool("h", false, "display help information")
	helpFlagLong = flag.Bool("help", false, "display help information")
	logLevel     = flag.String("log-level", "info", "set log level (debug, info, warn, error)")
)

func main() {
	flag.Parse()

	if *helpFlag || *helpFlagLong {
		printHelp()
		os.Exit(0)
	}

	logging.Init(*logLevel)
	logger := logging.With("component", "main")

	logger.Info("starting schoty", "version", "0.1.0")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		select {
		case sig := <-sigChan:
			logger.Info("received signal, initiating graceful shutdown", "signal", sig.String())
			cancel()
		case <-ctx.Done():
		}
	}()

	model := ui.NewModel()
	p := tea.NewProgram(&model, tea.WithContext(ctx))

	if err := p.Start(); err != nil {
		logging.Error("error starting schoty", "error", err.Error())
		os.Exit(1)
	}

	logger.Info("schoty exited successfully")

	signal.Stop(sigChan)
	wg.Wait()
}

func printHelp() {
	fmt.Println(`Schoty - AI Subscription Usage Monitor

Usage: schoty [options]

Options:
  -h, --help      display this help information
  --log-level     set log level: debug, info, warn, error (default: info)

Keyboard Shortcuts (when running):
  q, ctrl+c        quit application
  r                refresh all data
  c                toggle config view
  ?                show keyboard shortcuts help
  tab              cycle through tabs
  1-8              quick jump to provider
  up/down          navigate list
  enter            open detail view
  esc              go back / close

For more information, see: https://github.com/ChrisUFO/Schoty`)
}
