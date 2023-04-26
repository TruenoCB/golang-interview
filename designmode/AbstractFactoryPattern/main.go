package main

import (
	"abstractfactory/factory"
	"fmt"
)

func main() {
	xFactory, _ := factory.GetFactory("xiaomi")
	hFactory, _ := factory.GetFactory("huawei")
	xpn := xFactory.MakePhone().GetName()
	hpn := hFactory.MakePhone().GetName()
	fmt.Printf("%v,%v\n", xpn, hpn)
}
