package main

import "fmt"

func main() {
	nums := []int{1234, 34345, 345, 241, 756, 77, 4, 5234, 346, 35, 25, 3, 58, 33, 423, 41, 3567, 35, 26, 1, 6, 2, 3, 5, 67, 8, 4, 8, 9, 4356}
	shellSort(nums)
	fmt.Printf("%v\n", nums)
}

func shellSort(nums []int) {
	gap := len(nums)
	for gap > 1 {
		gap = gap/3 + 1
		for i := 0; i < len(nums)-gap; i++ {
			end := i
			x := nums[end+gap]
			for end >= 0 {
				if nums[end] > x {
					nums[end+gap] = nums[end]
					end -= gap
				} else {
					break
				}
			}
			nums[end+gap] = x
		}
	}
}
