package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if item, ok := c.items[key]; ok {
		item.Value.(*cacheItem).value = value
		c.queue.MoveToFront(item)
		return true
	}
	item := c.queue.PushFront(&cacheItem{key, value})

	c.items[key] = item

	if c.queue.Len() > c.capacity {
		if item := c.queue.Back(); item != nil {
			c.queue.Remove(item)
			delete(c.items, item.Value.(*cacheItem).key)
		}
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		return item.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
