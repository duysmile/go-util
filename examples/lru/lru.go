package main

import (
	"fmt"
	"github.com/duysmile/go-util/lru"
)

func main() {
	lruCache := lru.NewLRU(2)

	lruCache.Add("1", 1)
	lruCache.Add("2", 2)

	key := "2"
	v, _ := lruCache.Get(key)
	fmt.Printf("Get a cache hit [key]: %s, [value]: %v\n", key, v)

	lruCache.Add("3", 3)
	key = "1"
	v, _ = lruCache.Get(key)
	fmt.Printf("Cannot get cache [key]: %s, [value]: %v because cache reach max entries\n", key, v)
}
