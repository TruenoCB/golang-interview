package main

import (
	"lru/lru"
)

func main() {
	lru := lru.New(6)
	for i := 0; i < 100000; i++ {
		lru.Put(i, i+1)
	}
	for i := 0; i < 3; i++ {
		lru.Put(i, i+1)
	}
	lru.PrintMap()
}
