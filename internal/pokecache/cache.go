package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct{
	createdAt 	time.Time
	data 		[]byte
}

type Cache struct{
	mp		map[string]cacheEntry
	mx 		*sync.Mutex
	dur		time.Duration
}

func NewCache(duration time.Duration) Cache{
	toReturn :=  Cache {
		mp: make(map[string]cacheEntry),
		mx: &sync.Mutex{},
		dur: duration,
	}
	go toReturn.reapLoop()
	return toReturn
}

func (c *Cache) Add(key string, val []byte){
	entry := cacheEntry{
		createdAt: time.Now(),
		data: val,
	}
	c.mx.Lock()
	defer c.mx.Unlock()
	c.mp[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mx.Lock()
	defer c.mx.Unlock()
	val, found := c.mp[key];
	if found{
		return val.data, true
	}
	return nil, false
}

func (c *Cache) reapLoop(){
	ticker := time.Tick(c.dur)
	for {
		<-ticker
		c.reap()
	}
}

func (c *Cache) reap(){
	for key, val := range c.mp{
		passed := time.Since(val.createdAt)
		if passed > c.dur{
			delete(c.mp, key)
		}
	}
}
