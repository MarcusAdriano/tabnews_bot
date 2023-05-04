package cache

import (
	"sync"
	"time"
)

type Cache struct {
	lock sync.RWMutex
	items map[string]Item
}

type Item struct {
	Object     interface{}
	Expiration int64
}

func (c *Cache) Set(k string, v interface{}, d time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items[k] = Item{
		Object:     v,
		Expiration: time.Now().Add(d).UnixNano(),
	}
}

func (c *Cache) Get(k string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	item, found := c.items[k]
	if !found {
		return nil, false
	}
	if item.Expiration < time.Now().UnixNano() {
		return nil, false
	}
	return item.Object, true
}

func (c *Cache) Delete(k string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.items, k)
}

func (c *Cache) IsExpired(k string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	item, found := c.items[k]
	if !found {
		return false
	}
	return item.Expiration < time.Now().UnixNano()
}

func (c *Cache) DeleteExpired() {
	c.lock.Lock()
	defer c.lock.Unlock()
	for k, v := range c.items {
		if v.Expiration < time.Now().UnixNano() {
			delete(c.items, k)
		}
	}
}

func NewCache() *Cache {
	items := make(map[string]Item)
	cache := Cache{
		items: items,
	}

	return &cache
}