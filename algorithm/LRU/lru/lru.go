package lru

import (
	"container/list"
	"fmt"
)

type LRU struct {
	size      int
	innerList *list.List
	InnerMap  map[int]*list.Element
}

func New(size int) *LRU {
	return &LRU{size, list.New(), make(map[int]*list.Element)}
}

type entry struct {
	key   int
	value int
}

func (lru *LRU) Get(key int) (int, bool) {
	if e, ok := lru.InnerMap[key]; ok {
		lru.innerList.MoveToFront(e)
		return e.Value.(*entry).value, true
	}
	return -1, false
}

func (lru *LRU) Put(key, value int) {
	if e, ok := lru.InnerMap[key]; ok {
		e.Value.(*entry).value = value
		lru.size++
		lru.innerList.MoveToFront(e)
	} else {
		newentry := &entry{key, value}
		lru.InnerMap[key] = lru.innerList.PushFront(newentry)
		if lru.innerList.Len() > lru.size {
			back := lru.innerList.Back()
			lru.innerList.Remove(back)
			delete(lru.InnerMap, back.Value.(*entry).key)
		}
	}
}

func (lru *LRU) PrintMap() {
	for _, v := range lru.InnerMap {
		fmt.Printf("%v\n", v.Value.(*entry))
	}
}
