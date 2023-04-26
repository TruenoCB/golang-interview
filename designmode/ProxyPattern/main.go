package main

import (
	"fmt"
	"proxypattern/proxy"
)

func main() {
	myserver := proxy.GetMyServer(2)
	state, con := myserver.HandleRequest("Get", "1")
	fmt.Printf("%v,%v\n", state, con)
	state, con = myserver.HandleRequest("Get", "1")
	fmt.Printf("%v,%v\n", state, con)
	state, con = myserver.HandleRequest("Get", "1")
	fmt.Printf("%v,%v\n", state, con)
	state, con = myserver.HandleRequest("Get", "1")
	fmt.Printf("%v,%v\n", state, con)
	state, con = myserver.HandleRequest("Post", "2")
	fmt.Printf("%v,%v\n", state, con)
	state, con = myserver.HandleRequest("Get", "3")
	fmt.Printf("%v,%v\n", state, con)
}
