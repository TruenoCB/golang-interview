package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{1, 2, 3, 23, 4, 234, 23, 45, 4, 4234, 633, 45234, 5, 346, 25, 76, 5, 235, 34342, 414, 25, 4, 25, 32456, 252, 342, 5, 4234, 6, 547, 4, 6767, 969, 05, 473, 67, 57, 25, 364736, 34}
	sortArray(nums)
	fmt.Printf("%v\n", nums)

}

func sortArray(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start, end int) {
	if start < end {
		n := partition(nums, start, end)
		quickSort(nums, start, n-1)
		quickSort(nums, n+1, end)
	}
}

func partition(nums []int, start, end int) int {
	p1 := start - 1
	p2 := start
	randIndex := rand.Intn(end-start+1) + start
	swap(nums, randIndex, end)
	n := nums[end]
	for p2 < end {
		if nums[p2] < n {
			p1++
			swap(nums, p1, p2)
		}
		p2++
	}
	p1++
	swap(nums, p1, p2)
	return p1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
