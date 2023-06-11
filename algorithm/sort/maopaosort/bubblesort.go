package main

import "fmt"

func main() {
	nums := []int{2, 1234, 235, 43, 45, 6, 6, 8, 345, 78, 3467, 34, 78, 9, 67, 8, 346, 3245, 675, 57, 23, 52, 5, 473, 6, 56, 8}
	newBubbleSort(nums)
	fmt.Printf("%v\n", nums)
}

func newBubbleSort(nums []int) {
	lenth := len(nums) - 1
	tempPosition := 0 //优化内循环次数，只需要循环到最后一次交换的位置即可
	for i := 0; i < len(nums)-1; i++ {
		flag := 1 //优化标记位  如果某次内循环没有发生比较，说明已经排序完毕，可以直接退出。
		for j := 0; j < lenth; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = 0
				tempPosition = j
			}
		}
		if flag == 1 {
			return
		}
		lenth = tempPosition
	}
}
