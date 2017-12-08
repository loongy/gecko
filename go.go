package gecko

// Go runs a function on a goroutine. It returns a channel to which the result
// will be written, and a channel to which the error will be written.
func Go(f func() (interface{}, error)) (chan interface{}, chan error) {
	ret := make(chan interface{})
	err := make(chan error)
	go func() {
		r, e := f()
		if e != nil {
			err <- e
			return
		}
		ret <- r
	}()
	return ret, err
}
