// Package streamer fans-in container log streams from many pods into a
// single Printer.
package streamer

import (
	"bufio"
	"context"
	"fmt"
	"regexp"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"example.com/golab/17-capstone/internal/printer"
)

// Config holds the orchestrator's inputs.
type Config struct {
	Clientset kubernetes.Interface
	Namespace string
	Selector  string
	Container string         // "" = all containers
	Since     time.Duration  // 0 = from start
	Grep      *regexp.Regexp // nil = match all
	Printer   *printer.Printer
}

// Run lists matching pods, starts a streamer per (pod, container), and
// returns when ctx is canceled or all streams close.
//
// Implementation roadmap:
//
//  1. List pods with the selector.
//  2. Build a job queue: one entry per (pod-name, container-name) of READY pods.
//  3. Start a sync.WaitGroup; for each job, go func { defer wg.Done(); streamPod(...) }.
//  4. wg.Wait(); return ctx.Err() if canceled.
//
// Stage-5 stretch (after Stage 4 works): also start a watcher goroutine that
// receives Pod ADD/MODIFY events and starts streamers for any newly-ready
// pod that isn't already being tailed. Track tailed-set with a mutex.
//
// TODO: implement Run.
func Run(ctx context.Context, cfg Config) error {
	// TODO
	_ = corev1.PodReady
	_ = metav1.ListOptions{}
	return fmt.Errorf("Run not implemented yet")
}

// streamPod opens a log stream for one container and forwards each line
// through cfg.Printer until EOF, an error, or ctx is canceled.
//
// Use:
//
//	req := cfg.Clientset.CoreV1().Pods(cfg.Namespace).GetLogs(podName, &corev1.PodLogOptions{
//	    Container: container,
//	    Follow:    true,
//	    SinceSeconds: <pointer to int64 if cfg.Since > 0, else nil>,
//	})
//	stream, err := req.Stream(ctx)
//	... defer stream.Close() ...
//	scanner := bufio.NewScanner(stream)
//	for scanner.Scan() {
//	    text := scanner.Text()
//	    if cfg.Grep != nil && !cfg.Grep.MatchString(text) { continue }
//	    cfg.Printer.Print(printer.LogLine{Time: time.Now(), Pod: podName, Container: container, Text: text})
//	}
//
// Return nil on EOF, ctx.Err() on cancel, otherwise wrap the error.
//
// TODO: implement streamPod.
func streamPod(ctx context.Context, cfg Config, podName, container string) error {
	// TODO
	_ = bufio.NewScanner
	return nil
}

// isReady returns true if the pod's PodReady condition is True.
func isReady(p corev1.Pod) bool {
	for _, c := range p.Status.Conditions {
		if c.Type == corev1.PodReady && c.Status == corev1.ConditionTrue {
			return true
		}
	}
	return false
}

// containersOf returns the names of the containers in p, optionally filtered
// to a single container name. If filter is empty, returns all containers.
func containersOf(p corev1.Pod, filter string) []string {
	var out []string
	for _, c := range p.Spec.Containers {
		if filter == "" || c.Name == filter {
			out = append(out, c.Name)
		}
	}
	return out
}

// trackedSet is a goroutine-safe set of "pod/container" keys, used by Run's
// watcher to avoid double-tailing. (Wired up in Stage 5.)
type trackedSet struct {
	mu sync.Mutex
	m  map[string]struct{}
}

func newTrackedSet() *trackedSet { return &trackedSet{m: map[string]struct{}{}} }

func (t *trackedSet) add(key string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if _, ok := t.m[key]; ok {
		return false
	}
	t.m[key] = struct{}{}
	return true
}

func (t *trackedSet) remove(key string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.m, key)
}
