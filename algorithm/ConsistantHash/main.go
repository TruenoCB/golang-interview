package main

import (
	"conhash/hash"
	"fmt"
)

func main() {
	hashring := hash.New(50)
	hashring.Add("1节点")
	hashring.Add("2节点")
	hashring.Add("3节点")
	hashring.Add("4节点")
	hashring.Add("5节点")
	hashring.Add("6节点")
	for i := 0; i < 1000; i++ {
		fmt.Printf("key%v选择了%v\n", i, hashring.Get(string(i)))
	}
}
