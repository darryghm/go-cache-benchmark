package main

import (
	"github.com/bluele/gcache"
)

type GcacheLRU struct {
	cache gcache.Cache
}

func NewGcacheLRU(size int) Cache {
	cache := gcache.New(size).LRU().Build()
	return &GcacheLRU{
		cache: cache,
	}
}

func (c *GcacheLRU) Name() string {
	return "gcache"
}

func (c *GcacheLRU) Set(key string) {
	c.cache.Set(key, key)
}

func (c *GcacheLRU) Get(key string) bool {
	_, err := c.cache.Get(key)
	return err == nil
}

func (c *GcacheLRU) Close() {
}