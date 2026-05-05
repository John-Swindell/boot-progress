// Command kubetail tails container logs from every pod matching a label
// selector in a namespace, multiplexed and color-coded.
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"example.com/golab/17-capstone/internal/k8s"
	"example.com/golab/17-capstone/internal/printer"
	"example.com/golab/17-capstone/internal/streamer"
)

type opts struct {
	namespace  string
	selector   string
	container  string
	kubeconfig string
	since      time.Duration
	grep       string
	json       bool
	noColor    bool
}

func parseFlags() (opts, error) {
	var o opts
	fs := flag.NewFlagSet("kubetail", flag.ContinueOnError)
	fs.StringVar(&o.namespace, "n", "default", "namespace")
	fs.StringVar(&o.selector, "l", "", "label selector (e.g. app=loggy)")
	fs.StringVar(&o.container, "container", "", "container name (default: all containers)")
	fs.StringVar(&o.kubeconfig, "kubeconfig", "", "path to kubeconfig (default: $KUBECONFIG or ~/.kube/config)")
	fs.DurationVar(&o.since, "since", 0, "show logs newer than this duration (e.g. 5m); 0 = from start of pod")
	fs.StringVar(&o.grep, "grep", "", "only print lines matching this regex")
	fs.BoolVar(&o.json, "json", false, "emit JSON lines instead of pretty format")
	fs.BoolVar(&o.noColor, "no-color", false, "disable ANSI color")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return o, err
	}
	return o, nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	o, err := parseFlags()
	if err != nil {
		os.Exit(2)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, o); err != nil {
		slog.Error("exit", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, o opts) error {
	cs, err := k8s.NewClientset(o.kubeconfig)
	if err != nil {
		return fmt.Errorf("kube client: %w", err)
	}

	var grepRE *regexp.Regexp
	if o.grep != "" {
		grepRE, err = regexp.Compile(o.grep)
		if err != nil {
			return fmt.Errorf("invalid -grep regex: %w", err)
		}
	}

	p := printer.New(os.Stdout, printer.Options{
		JSON:    o.json,
		NoColor: o.noColor,
	})

	return streamer.Run(ctx, streamer.Config{
		Clientset: cs,
		Namespace: o.namespace,
		Selector:  o.selector,
		Container: o.container,
		Since:     o.since,
		Grep:      grepRE,
		Printer:   p,
	})
}
