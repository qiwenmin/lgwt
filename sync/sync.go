package sync

import "sync"

// Counter - a basic counter.
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter creates a new counter.
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increases the counter.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

// Value returns the counter value.
func (c *Counter) Value() int {
	return c.value
}
