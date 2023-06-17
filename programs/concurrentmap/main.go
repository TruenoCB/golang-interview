package main

import (
	"cmap/concurrentmap"
	"fmt"
	"time"
)

func main() {
	concurrentmap := concurrentmap.NewConMap()
	nowtime := time.Now()
	go func() {
		v, err := concurrentmap.Get(1, time.Second)
		if err == nil {
			fmt.Printf("%v,%v\n", v, time.Now().Sub(nowtime))
		} else {
			fmt.Println(err.Error())
		}
	}()
	go func() {
		v, err := concurrentmap.Get(1, 5*time.Second)
		if err == nil {
			fmt.Printf("%v,%v\n", v, time.Now().Sub(nowtime))
		} else {
			fmt.Println(err.Error())
		}
	}()
	time.Sleep(2 * time.Second)
	go concurrentmap.Put(1, 1)
	time.Sleep(10 * time.Second)
}
