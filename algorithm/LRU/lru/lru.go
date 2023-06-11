package lru

import (
	"container/list"
	"fmt"
)

type LRU struct {
	size      int
	innerList *list.List
	innerMap  map[int]*list.Element
}

func New(size int) *LRU {
	return &LRU{size, list.New(), make(map[int]*list.Element)}
}

type entry struct {
	key   int
	value int
}

func (lru *LRU) Get(key int) (int, bool) {
	if e, ok := lru.innerMap[key]; ok {
		lru.innerList.MoveToFront(e)
		return e.Value.(*entry).value, true
	}
	return -1, false
}

func (lru *LRU) Put(key, value int) {
	if e, ok := lru.innerMap[key]; ok {
		e.Value.(*entry).value = value
		lru.innerList.MoveToFront(e)
	} else {
		newentry := &entry{key, value}
		lru.innerMap[key] = lru.innerList.PushFront(newentry)
		if lru.innerList.Len() > lru.size {
			back := lru.innerList.Back()
			lru.innerList.Remove(back)
			delete(lru.innerMap, back.Value.(*entry).key)
		}
	}
}

func (lru *LRU) PrintMap() {
	for _, v := range lru.innerMap {
		fmt.Printf("%v\n", v.Value.(*entry))
	}
}
