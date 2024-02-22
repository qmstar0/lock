package lock

import "sync"

type Closer interface {
	Closed() bool
	Closing() <-chan struct{}
	CloseAndRunOnce(onClose func() error) error
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
func (c *closer) CloseAndRunOnce(onClose func() error) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.closed {
		return nil
	}
	c.closed = true
	close(c.closing)
	if onClose != nil {
		err := onClose()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *closer) Reset() {
	c.closed = false
	c.closing = make(chan struct{})
}
