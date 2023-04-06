package main

import "fmt"

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

func main() {
	//调用迭代
	nums1 := []int{123, 1, 24, 211, 3, 12, 5, 76, 84, 95, 345, 48, 35, 135, 6, 42312, 52, 68, 985, 6, 54, 3, 3245, 6, 37, 5, 5, 58, 8, 56, 52, 52, 5}
	mergeSortIteration(nums1)
	fmt.Printf("%v\n", nums1)

	//调用递归
	nums2 := []int{123, 1, 24, 211, 3, 12, 5, 76, 84, 95, 345, 48, 35, 135, 6, 42312, 52, 68, 985, 6, 54, 3, 3245, 6, 37, 5, 5, 58, 8, 56, 52, 52, 5}
	dst := []int{}
	dst = append(dst, nums2...)
	mergeSortRecursion(nums2, dst, 0, len(dst))
	fmt.Printf("%v\n", dst)

}

// 迭代实现归并排序
func mergeSortIteration(nums []int) {
	lenth := len(nums)
	src := nums //src源数组 dst目标数组 操作后要交换两个数组
	dst := make([]int, lenth)
	for seg := 1; seg < lenth; seg = seg << 1 {
		for start := 0; start < lenth; start += seg * 2 {
			mid := min(start+seg, lenth)
			end := min(start+seg*2, lenth)
			i, j, k := start, mid, start //注意start，end左闭右开
			for i < mid || j < end {
				if j == end || (i < mid && (src[i] < src[j])) {
					dst[k] = src[i]
					i++
				} else {
					dst[k] = src[j]
					j++
				}
				k++
			}
		}
		src, dst = dst, src
	}
	nums = src
}

// 递归实现归并排序
func mergeSortRecursion(src, dst []int, start, end int) { //注意start，end左闭右开
	if start+1 >= end {
		return
	}
	mid := (start + end) / 2
	mergeSortRecursion(dst, src, start, mid) //注意交换src，dst    这一步的目的是为了将src的start, mid变成顺序的，所以要交换dist和mid
	mergeSortRecursion(dst, src, mid, end)
	i, j, k := start, mid, start
	for i < mid || j < end {
		if j == end || (i < mid && (src[i] < src[j])) {
			dst[k] = src[i]
			i++
		} else {
			dst[k] = src[j]
			j++
		}
		k++
	}
}
