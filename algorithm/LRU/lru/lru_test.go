package lru

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	lru := New(666)
	for i := 0; i < 100000; i++ {
		lru.Put(i, i+1)
	}
	fmt.Printf("%v\n", lru.innerList)
}
