package cache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu                 sync.Mutex
	originImagePath    string
	croppedImagePath   string
	isClearCacheImages bool
	capacity           int
	queue              List
	items              map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int, originImagePath, croppedImagePath string, isClearCacheImages bool) Cache {
	return &lruCache{
		capacity:           capacity,
		originImagePath:    originImagePath,
		croppedImagePath:   croppedImagePath,
		isClearCacheImages: isClearCacheImages,
		queue:              NewList(),
		items:              make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	newItem := cacheItem{value: value, key: key}
	item, ok := c.items[key]

	if ok {
		c.queue.MoveToFront(item)
		item.Value = newItem

		return true
	}

	current := c.queue.PushFront(newItem)
	c.items[key] = current

	if c.queue.Len() > c.capacity {
		last := c.queue.Back()
		c.queue.Remove(last)
		key := last.Value.(cacheItem).key
		delete(c.items, key)

		if c.isClearCacheImages {
			RemoveCacheImages(c.originImagePath, c.croppedImagePath, string(key))
		}

		return false
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.items[key]

	if ok {
		c.queue.MoveToFront(item)
		return c.items[key].Value.(cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}
