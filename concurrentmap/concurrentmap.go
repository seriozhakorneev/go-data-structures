package concurrentmap

import (
	"sync"
)

// CMap - concurrent-safety map.
type CMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
	// generic zero value
	zero V
}

// New - returns new CMap exemplar.
func New[K comparable, V any]() *CMap[K, V] {
	return &CMap[K, V]{
		mu: sync.RWMutex{},
		m:  make(map[K]V),
	}
}

// Insert - creates new entry with provided key, value.
func (c *CMap[K, V]) Insert(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = value
}

// Get - returns value lying at provided key and true,
// if entry with key not exist, returns zero value  and false.
func (c *CMap[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if value, ok := c.m[key]; ok {
		return value, true
	}

	return c.zero, false
}

// Update - returns true, if value for key are set, else returns false.
func (c *CMap[K, V]) Update(key K, value V) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.m[key]; ok {
		c.m[key] = value
		return true
	}

	return false
}

// Delete - returns true, if entry by provided key are deleted,
// else return false.
func (c *CMap[K, V]) Delete(key K) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.m[key]; ok {
		delete(c.m, key)
		return true
	}

	return false
}

// getMap - returns underlying map m
func (c *CMap[K, V]) getMap() map[K]V {
	return c.m
}
