package lock

import "sync"

type Closer struct {
	Closed  bool
	closing chan struct{}
	sync.Mutex
}

func (c *Closer) Closing() <-chan struct{} {
	if c.closing == nil {
		c.closing = make(chan struct{})
	}
	return c.closing
}
