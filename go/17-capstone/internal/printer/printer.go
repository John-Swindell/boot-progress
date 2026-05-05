// Package printer formats LogLines for human or JSON output.
package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"
)

// LogLine is one structured log line emitted by a streamer goroutine.
type LogLine struct {
	Time      time.Time
	Pod       string
	Container string
	Text      string
}

// Options controls Printer rendering.
type Options struct {
	JSON    bool
	NoColor bool
}

// Printer is safe for concurrent use; multiple streamer goroutines can call
// Print without producing torn lines.
type Printer struct {
	mu   sync.Mutex
	w    io.Writer
	opts Options
}

// New returns a Printer writing to w.
func New(w io.Writer, opts Options) *Printer {
	return &Printer{w: w, opts: opts}
}

// Print formats and writes one log line.
func (p *Printer) Print(line LogLine) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	out, err := p.format(line)
	if err != nil {
		return err
	}
	_, err = p.w.Write(out)
	return err
}

// format renders a LogLine according to opts.
//
// TODO: implement.
//
// JSON mode (opts.JSON == true):
//
//	one line of JSON ending in '\n', encoded with encoding/json:
//	{"ts": "<RFC3339Nano>", "pod": "<pod>", "container": "<container>", "line": "<text>"}
//
// Pretty mode (opts.JSON == false):
//
//	"<COLOR>[<pod>/<container>]<RESET> <text>\n"
//	where <COLOR> is colorFor(line.Pod+"/"+line.Container) UNLESS opts.NoColor,
//	in which case omit color escapes entirely.
//
// In both modes the trailing newline is part of the output. line.Text never
// contains a trailing newline (the streamer strips it).
//
// Hint:
//
//	json mode -> json.Marshal a small map / struct, append '\n'
//	pretty mode -> fmt.Sprintf
func (p *Printer) format(line LogLine) ([]byte, error) {
	// TODO: replace this stub
	_ = json.Marshal
	_ = fmt.Sprintf
	return nil, nil
}
