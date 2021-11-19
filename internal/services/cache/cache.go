package cache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) (bool, Key)
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) (isRe bool, removedKey Key) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newItem := cacheItem{value: value, key: key}
	item, ok := c.items[key]

	if ok {
		c.queue.MoveToFront(item)
		item.Value = newItem
		return true, ""
	}

	current := c.queue.PushFront(newItem)
	c.items[key] = current

	if c.queue.Len() > c.capacity {
		last := c.queue.Back()
		c.queue.Remove(last)
		key := last.Value.(cacheItem).key
		delete(c.items, key)

		return false, key
	}

	return false, ""
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()

	item, ok := c.items[key]

	if ok {
		c.queue.MoveToFront(item)
		defer c.mu.Unlock()
		return c.items[key].Value.(cacheItem).value, true
	}

	defer c.mu.Unlock()
	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}
