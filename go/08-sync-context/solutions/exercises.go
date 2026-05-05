package solutions

import (
	"context"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	n  int
}

func (s *SafeCounter) Inc() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.n++
}

func (s *SafeCounter) Value() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.n
}

func ParallelMap(xs []int, fn func(int) int) []int {
	out := make([]int, len(xs))
	var wg sync.WaitGroup
	for i, x := range xs {
		wg.Add(1)
		go func(i, x int) {
			defer wg.Done()
			out[i] = fn(x)
		}(i, x)
	}
	wg.Wait()
	return out
}

func Memoize(loader func() (string, error)) func() (string, error) {
	var (
		once sync.Once
		v    string
		err  error
	)
	return func() (string, error) {
		once.Do(func() { v, err = loader() })
		return v, err
	}
}

func Sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

type result struct {
	v   string
	err error
}

func Race(ctx context.Context, fns []func(context.Context) (string, error)) (string, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	results := make(chan result, len(fns))
	for _, fn := range fns {
		go func(fn func(context.Context) (string, error)) {
			v, err := fn(ctx)
			results <- result{v, err}
		}(fn)
	}
	var lastErr error
	for i := 0; i < len(fns); i++ {
		r := <-results
		if r.err == nil {
			return r.v, nil
		}
		lastErr = r.err
	}
	return "", lastErr
}
