package main

import (
    lru "github.com/hashicorp/golang-lru/v2/expirable"
)

type HashicorpExpLRU struct {
    v *lru.LRU[string, string]
}

func NewHashicorpExpLRU(size int) Cache {
    cache := lru.NewLRU[string, string](size, nil, 0)
    return &HashicorpExpLRU{
        v: cache,
    }
}

func (c *HashicorpExpLRU) Name() string {
    return "expirable-lru-hashicorp"
}

func (c *HashicorpExpLRU) Set(key string) {
    c.v.Add(key, key)
}

func (c *HashicorpExpLRU) Get(key string) bool {
    _, ok := c.v.Get(key)
    return ok
}

func (c *HashicorpExpLRU) Close() {}
