package gecko

import (
	"reflect"
	"runtime"
	"sync"
)

// An Generator defines the number of iterations that will be distributed over
// the available CPUs.
type Generator int

// Do the number of iterations, defined by the Iterator. Iterations will be
// distributed over the available CPUs with an approximately equal workload.
func (gen Generator) Do(f func(i int)) {
	var wg sync.WaitGroup
	defer wg.Wait()

	// Decide on how many parallel goroutines will be used, and the workload for
	// each.
	numCPUs := runtime.NumCPU()
	numPerCPU := int(gen)/numCPUs + 1
	wg.Add(numCPUs)

	// Run goroutines.
	for i := 0; i < numCPUs; i++ {
		go func(i int) {
			defer wg.Done()
			for j := i * numPerCPU; j < (i+1)*numPerCPU && j < int(gen); j++ {
				f(j)
			}
		}(i)
	}
}

// ForAll returns an Generator that spans arrays, slices, or an integer range.
func ForAll(gen interface{}) Generator {
	switch reflect.TypeOf(gen).Kind() {
	case reflect.Array:
		return Generator(reflect.ValueOf(gen).Len())
	case reflect.Slice:
		return Generator(reflect.ValueOf(gen).Len())
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return Generator(reflect.ValueOf(gen).Int())
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Generator(reflect.ValueOf(gen).Uint())
	default:
		// This will panic with a nice, consistent, error message about invalid
		// conversions.
		return gen.(Generator)
	}
}
