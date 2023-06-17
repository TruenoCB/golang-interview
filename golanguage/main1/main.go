package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg = sync.WaitGroup{}
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	wg.Add(3)
	go dog(ch1, ch2)
	go cat(ch2, ch3)
	go fish(ch3, ch1)
	ch1 <- struct{}{}
	wg.Wait()
	fmt.Println("打印结束")

}

func dog(ch1, ch2 chan struct{}) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-ch1
		fmt.Println("dog")
		select {
		case ch2 <- struct{}{}:
		default:
		}
	}
}

func cat(ch1, ch2 chan struct{}) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-ch1
		fmt.Println("cat")
		select {
		case ch2 <- struct{}{}:
		default:
		}
	}
}

func fish(ch1, ch2 chan struct{}) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-ch1
		fmt.Println("fish")
		select {
		case ch2 <- struct{}{}:
		default:
		}
	}
}
