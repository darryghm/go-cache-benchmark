package main

import (
	"github.com/coocood/freecache"
)

type FreeCacheLRU struct {
	cache *freecache.Cache
}

func NewFreeCacheLRU(size int) Cache {
	cache := freecache.NewCache(size)
	return &FreeCacheLRU{
		cache: cache,
	}
}

func (c *FreeCacheLRU) Name() string {
	return "freecache"
}

func (c *FreeCacheLRU) Set(key string) {
	c.cache.Set([]byte(key), []byte(key), 0)
}

func (c *FreeCacheLRU) Get(key string) bool {
	_, err := c.cache.Get([]byte(key))
	return err == nil
}

func (c *FreeCacheLRU) Close() {}