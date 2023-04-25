package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{23, 56, 34, 45, 23, 46, 34, 46, 34, 456, 65, 67, 45, 56, 2, 5, 32, 6, 8, 34, 7, 3, 8, 9, 23}
	countsort(nums)
	fmt.Printf("%v\n", nums)
}

func countsort(nums []int) {
	minn := math.MaxInt64
	maxn := math.MinInt64
	for _, n := range nums {
		minn = min(n, minn)
		maxn = max(n, maxn)
	}
	count := make([]int, maxn-minn+1)
	for _, n := range nums {
		count[n-minn]++
	}
	numsi := 0
	for i, _ := range count {
		for count[i] > 0 {
			nums[numsi] = i + minn
			numsi++
			count[i]--
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
