package lock

import "sync"

type Closer interface {
	Closed() bool
	Closing() <-chan struct{}
	LockedRun(fn func() error) error
	Reset()
}
type closer struct {
	closed  bool
	closing chan struct{}
	lock    sync.Mutex
}

func NewCloser() Closer {
	return &closer{
		closed:  false,
		closing: make(chan struct{}),
		lock:    sync.Mutex{},
	}
}

func (c *closer) Closed() bool {
	return c.closed
}

func (c *closer) Closing() <-chan struct{} {
	return c.closing
}
func (c *closer) LockedRun(fn func() error) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return fn()
}

func (c *closer) Reset() {
	c.closed = false
	c.closing = make(chan struct{})
}
