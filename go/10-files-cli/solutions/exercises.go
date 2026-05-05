package solutions

import (
	"bufio"
	"flag"
	"io"
	"os"
	"strings"
)

func WordCount(r io.Reader) (lines, words, bytes int, err error) {
	data, e := io.ReadAll(r)
	if e != nil {
		return 0, 0, 0, e
	}
	bytes = len(data)
	if bytes == 0 {
		return 0, 0, 0, nil
	}
	s := bufio.NewScanner(strings.NewReader(string(data)))
	for s.Scan() {
		lines++
		words += len(strings.Fields(s.Text()))
	}
	if e := s.Err(); e != nil {
		return 0, 0, 0, e
	}
	return lines, words, bytes, nil
}

func WriteAtomically(path string, data []byte) error {
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		os.Remove(tmp)
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		os.Remove(tmp)
		return err
	}
	return nil
}

type AppConfig struct {
	Addr    string
	Workers int
	Debug   bool
}

func ParseFlags(args []string) (AppConfig, error) {
	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // don't dump usage during tests
	var cfg AppConfig
	fs.StringVar(&cfg.Addr, "addr", ":8080", "listen address")
	fs.IntVar(&cfg.Workers, "workers", 4, "concurrent workers")
	fs.BoolVar(&cfg.Debug, "debug", false, "enable debug logging")
	if err := fs.Parse(args); err != nil {
		return AppConfig{}, err
	}
	return cfg, nil
}
