package main

import "fmt"

func main() {
	nums := []int{3, 4, 63, 4, 52, 42, 5364, 55, 2, 54, 63, 2, 51, 5, 7, 5878, 18, 2}
	heapsort(nums)
	fmt.Printf("%v\n", nums)
}

func heapsort(nums []int) {
	for i := len(nums); i > 1; i-- {
		heapbuild(nums, i)
		swap(&nums[0], &nums[i-1])
	}

}

func swap(i, j *int) {
	*i, *j = *j, *i
}

func heapbuild(nums []int, len int) {
	parent := len/2 - 1
	for parent >= 0 {
		heapify(nums, parent, len)
		parent--
	}

}

func heapify(nums []int, parent int, len int) {
	max := parent
	lson := parent*2 + 1
	rson := parent*2 + 2
	if lson < len && nums[lson] > nums[max] {
		max = lson
	}
	if rson < len && nums[rson] > nums[max] {
		max = rson
	}
	if max != parent {
		swap(&nums[max], &nums[parent])
		heapify(nums, max, len)
	}

}
