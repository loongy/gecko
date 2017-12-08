package gecko

import (
	"errors"
	"testing"
)

func TestGo(t *testing.T) {
	// Goroutine.
	ret, err := Go(func() (interface{}, error) {
		return 420, nil
	})
	select {
	// A valid result received first.
	case r := <-ret:
		if r != 420 {
			t.Fatalf("unexpected result %v", r)
		}
	// An erro was received first.
	case e := <-err:
		t.Fatal(e)
	}
}

func TestGoError(t *testing.T) {
	// Goroutine.
	ret, err := Go(func() (interface{}, error) {
		return nil, errors.New("")
	})
	defer close(ret)
	defer close(err)

	select {
	// A valid result received first.
	case r := <-ret:
		t.Fatalf("unexpected result %v", r)
	// An erro was received first.
	case <-err:
		break
	}
}
