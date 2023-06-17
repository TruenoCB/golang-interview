package concurrentmap

import (
	"context"
	"sync"
	"time"
)

type concurrentmap struct {
	sync.Mutex
	imap map[int]int
	k2c  map[int]*mychan
}

type mychan struct {
	sync.Once
	ichan chan struct{}
}

func NewConMap() *concurrentmap {
	return &concurrentmap{imap: make(map[int]int), k2c: make(map[int]*mychan)}
}

func newmc() *mychan {
	return &mychan{ichan: make(chan struct{})}
}

func (mc *mychan) close() {
	mc.Do(func() {
		close(mc.ichan)
	})

}

func (c *concurrentmap) Put(k, v int) {
	c.Lock()
	defer c.Unlock()
	c.imap[k] = v
	if ch, ok := c.k2c[k]; ok {
		ch.close()
	}
}

func (c *concurrentmap) Get(k int, maxwaittime time.Duration) (int, error) {
	c.Lock()
	if v, ok := c.imap[k]; ok {
		c.Unlock()
		return v, nil
	}
	if _, ok := c.k2c[k]; !ok {
		c.k2c[k] = newmc()
	}

	tctx, cancel := context.WithTimeout(context.Background(), maxwaittime)
	defer cancel()
	c.Unlock()
	select {
	case <-tctx.Done():
		return -1, tctx.Err()
	case <-c.k2c[k].ichan:
	}
	c.Lock()
	defer c.Unlock()
	return c.imap[k], nil
}
