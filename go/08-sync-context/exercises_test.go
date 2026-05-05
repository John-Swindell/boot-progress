package syncctx

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestSafeCounter(t *testing.T) {
	var c SafeCounter
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	if c.Value() != 1000 {
		t.Errorf("Value = %d, want 1000", c.Value())
	}
}

func TestParallelMap(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	got := ParallelMap(xs, func(x int) int { return x * 10 })
	want := []int{10, 20, 30, 40, 50}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestMemoize(t *testing.T) {
	calls := 0
	var mu sync.Mutex
	loader := func() (string, error) {
		mu.Lock()
		defer mu.Unlock()
		calls++
		return "hello", nil
	}
	memo := Memoize(loader)

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v, err := memo()
			if err != nil || v != "hello" {
				t.Errorf("got (%q, %v)", v, err)
			}
		}()
	}
	wg.Wait()

	mu.Lock()
	defer mu.Unlock()
	if calls != 1 {
		t.Errorf("loader called %d times, want 1", calls)
	}
}

func TestSleep(t *testing.T) {
	t.Run("elapses normally", func(t *testing.T) {
		ctx := context.Background()
		start := time.Now()
		if err := Sleep(ctx, 20*time.Millisecond); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
		if time.Since(start) < 15*time.Millisecond {
			t.Errorf("returned too early")
		}
	})
	t.Run("respects cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(10 * time.Millisecond)
			cancel()
		}()
		err := Sleep(ctx, time.Second)
		if !errors.Is(err, context.Canceled) {
			t.Errorf("err = %v, want context.Canceled", err)
		}
	})
}

func TestRace(t *testing.T) {
	t.Run("first winner", func(t *testing.T) {
		fns := []func(context.Context) (string, error){
			func(ctx context.Context) (string, error) {
				select {
				case <-time.After(50 * time.Millisecond):
					return "slow", nil
				case <-ctx.Done():
					return "", ctx.Err()
				}
			},
			func(ctx context.Context) (string, error) {
				return "fast", nil
			},
		}
		got, err := Race(context.Background(), fns)
		if err != nil || got != "fast" {
			t.Errorf("got (%q, %v), want (fast, nil)", got, err)
		}
	})
	t.Run("all fail", func(t *testing.T) {
		boom := errors.New("boom")
		fns := []func(context.Context) (string, error){
			func(ctx context.Context) (string, error) { return "", boom },
			func(ctx context.Context) (string, error) { return "", boom },
		}
		_, err := Race(context.Background(), fns)
		if !errors.Is(err, boom) {
			t.Errorf("err = %v, want boom", err)
		}
	})
}
