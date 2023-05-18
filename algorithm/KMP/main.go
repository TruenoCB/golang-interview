package main

import (
	"fmt"
	"kmp"
)

func main() {
	fmt.Printf("%v\n", kmp.GetNext("aabaaf"))
	fmt.Printf("%v\n", kmp.StrStr("aabaabaafaabaabaafaabaabaafasfdaabaabaaf", "aabaaf"))
	fmt.Printf("%v\n", kmp.FirstStrStr("aabaabaaf", "aabaaf"))
}
