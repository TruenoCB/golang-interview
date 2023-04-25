package main

import (
	"fmt"
	"singleton/pkg"
)

func main() {
	for i := 0; i < 5; i++ {
		go pkg.GetInstance()
	}
	fmt.Scanln()
}
