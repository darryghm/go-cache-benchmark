package main

import (
	"runtime"
	"time"
)

type Benchmark struct {
	Generator
	N int
}

func main() {
	cacheSize := []int{1e3, 10e3, 100e3, 1e6}
	multiplier := []int{10, 100, 1000}
	newCache := []NewCacheFunc{
		NewGcacheLRU,
		NewHashicorpExpLRU,
		NewHashicorpLRU,
		NewFreeCacheLRU,
		NewTinyLFU,
		NewClockPro,
		NewARC,
		NewRistretto,
		NewDirectCache,
		NewTwoQueue,
		NewGroupCacheLRU,
		NewS4LRU,
		NewSLRU,
		NewWTFCache,
	}
	newGen := []NewGeneratorFunc{
		NewScrambledZipfian,
		// NewHotspot,
		// NewUniform,
	}

	var results []*BenchmarkResult

	for _, newGen := range newGen {
		for _, cacheSize := range cacheSize {
			for _, multiplier := range multiplier {
				numKey := cacheSize * multiplier

				for _, newCache := range newCache {
					result := run(newGen, cacheSize, numKey, newCache)
					results = append(results, result)
				}

				if len(results) > 0 {
					printResults(results)
					results = results[:0]
				}
			}
		}
	}
}

func run(newGen NewGeneratorFunc, cacheSize, numKey int, newCache NewCacheFunc) *BenchmarkResult {
	gen := newGen(numKey)
	b := &Benchmark{
		Generator: gen,
		N:         1e6,
	}

	alloc1 := memAlloc()
	cache := newCache(cacheSize)
	defer cache.Close()

	start := time.Now()
	hits, misses := bench(b, cache)
	dur := time.Since(start)

	alloc2 := memAlloc()

	return &BenchmarkResult{
		GenName:   gen.Name(),
		CacheName: cache.Name(),
		CacheSize: cacheSize,
		NumKey:    numKey,

		Hits:     hits,
		Misses:   misses,
		Duration: dur,
		Bytes:    int64(alloc2) - int64(alloc1),
	}
}

func bench(b *Benchmark, cache Cache) (hits, misses int) {
	for i := 0; i < b.N; i++ {
		value := b.Next()
		if cache.Get(value) {
			hits++
		} else {
			misses++
			cache.Set(value)
		}
	}

	return hits, misses
}

func memAlloc() uint64 {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}
