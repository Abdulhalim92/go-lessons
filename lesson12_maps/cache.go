package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.store[key]
	return value, exists
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
}

func caching() {
	cache := NewCache()
	cache.Set("foo", "bar")

	value, exists := cache.Get("foo")
	if exists {
		fmt.Println("Found value:", value)
	} else {
		fmt.Println("Value not found")
	}

	cache.Delete("foo")
	value, exists = cache.Get("foo")
	if exists {
		fmt.Println("Found value:", value)
	} else {
		fmt.Println("Value not found")
	}
}
