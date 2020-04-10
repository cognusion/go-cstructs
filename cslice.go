package cstructs

import (
	"sync"
)

// CSlice is a goro-safe interface{} slice
type CSlice struct {
	sync.RWMutex
	things []interface{}
}

// Cap returns the capacity of the underlying slice
func (c *CSlice) Cap() int {
	c.Lock()
	defer c.Unlock()
	return cap(c.things)
}

// Len returns the length of the underlying slice
func (c *CSlice) Len() int {
	c.Lock()
	defer c.Unlock()
	return len(c.things)
}

// InitSlice [re]initializes the underlying slice with the specified size,
// enabling safe Insert operations
func (c *CSlice) InitSlice(size int) {
	c.Lock()
	c.things = make([]interface{}, size)
	c.Unlock()
}

// Insert should be used only after InitSlice is used, and care is taken
// that 0 =< index < size, else panic
func (c *CSlice) Insert(index int, thing interface{}) {
	c.Lock()
	c.things[index] = thing
	c.Unlock()
}

// Get returns the value of the interface{} at the specific index.
// Care should be taken that 0 =< index < size, else panic
func (c *CSlice) Get(index int) interface{} {
	c.RLock()
	defer c.RUnlock()
	return c.things[index]
}

// GetSlice returns the underlying interface{} slice
func (c *CSlice) GetSlice() []interface{} {
	c.Lock()
	defer c.Unlock()
	return c.things
}

// Append s the specified item onto the underlying slice
func (c *CSlice) Append(thing interface{}) {
	c.Lock()
	c.things = append(c.things, thing)
	c.Unlock()
}
