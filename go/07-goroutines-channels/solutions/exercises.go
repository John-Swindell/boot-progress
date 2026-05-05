package solutions

import (
	"sync"
	"time"
)

func Generate(xs ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, x := range xs {
			out <- x
		}
	}()
	return out
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()
	return out
}

func Sum(in <-chan int) int {
	total := 0
	for v := range in {
		total += v
	}
	return total
}

func Merge(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, ch := range ins {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func FirstResponse(chans []<-chan int) (int, bool) {
	merged := Merge(chans...)
	v, ok := <-merged
	return v, ok
}

func WithTimeout(in <-chan int, d time.Duration) (int, bool) {
	select {
	case v, ok := <-in:
		if !ok {
			return 0, false
		}
		return v, true
	case <-time.After(d):
		return 0, false
	}
}
