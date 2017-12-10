package gecko

import (
	"testing"
	"time"
)

func TestForAll(t *testing.T) {
	xs := [10000]int{}
	for i := 0; i < len(xs); i++ {
		xs[i] = i
	}

	// Perform an increasingly heavy operation on each element in parallel and
	// measure the duration.
	start := time.Now()
	ForAll(xs).Do(func(i int) {
		w := 0
		for j := 0; j < i; j++ {
			w++
		}
	})
	duration := time.Now().Sub(start)

	// Perform an increasingly heavy operation on each element sequentially and
	// measure the duration.
	start = time.Now()
	for i := 0; i < len(xs); i++ {
		w := 0
		for j := 0; j < i; j++ {
			w++
		}
	}
	durationCmp := time.Now().Sub(start)

	// Compare the durations.
	if duration > durationCmp {
		t.Fatal("no performance improvements,", duration, "vs", durationCmp)
	}
}

func work(load int) {
	
}
