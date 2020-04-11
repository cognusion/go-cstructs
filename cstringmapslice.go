package cstructs

import (
	"sync"
)

// CStringMapSlice is a goro-safe map[string]interface{} slice
type CStringMapSlice struct {
	sync.RWMutex
	things []map[string]interface{}
}

// Cap returns the capacity of the underlying slice
func (c *CStringMapSlice) Cap() int {
	c.Lock()
	defer c.Unlock()
	return cap(c.things)
}

// Len returns the length of the underlying slice
func (c *CStringMapSlice) Len() int {
	c.Lock()
	defer c.Unlock()
	return len(c.things)
}

// InitSlice [re]initializes the underlying slice with the specified size,
// enabling safe Insert operations
func (c *CStringMapSlice) InitSlice(size int) {
	c.Lock()
	c.things = make([]map[string]interface{}, size)
	c.Unlock()
}

// Insert should be used only after InitSlice is used, and care is taken
// that 0 =< index < size, else panic
func (c *CStringMapSlice) Insert(index int, thing map[string]interface{}) {
	c.Lock()
	c.things[index] = thing
	c.Unlock()
}

// Get returns the value of the interface{} at the specific index.
// Care should be taken that 0 =< index < size, else panic
func (c *CStringMapSlice) Get(index int) map[string]interface{} {
	c.RLock()
	defer c.RUnlock()
	return c.things[index]
}

// GetSlice returns the underlying interface{} slice
func (c *CStringMapSlice) GetSlice() []map[string]interface{} {
	c.Lock()
	defer c.Unlock()
	return c.things
}

// Append s the specified item onto the underlying slice
func (c *CStringMapSlice) Append(thing map[string]interface{}) {
	c.Lock()
	c.things = append(c.things, thing)
	c.Unlock()
}
