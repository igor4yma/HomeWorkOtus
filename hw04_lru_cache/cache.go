package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache // Remove me after realization.

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

func (c *lruCache) Set(key Key, value interface{}) bool {
	item := cacheItem{key: key, value: value}
	if listItem, ok := c.items[key]; ok {
		c.queue.MoveToFront(listItem)
	} else {
		listItem := c.queue.PushFront(item)
		c.items[key] = listItem
	}
	return false
}
func (c *lruCache) Get(key Key) (interface{}, bool) {
	return nil, false
}
func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem)
	c.queue = NewList()

}
