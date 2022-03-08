/*
 * @Author: Alexleslie
 * @Date: 2022-03-08 15:22:04
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-08 19:45:37
 * @FilePath: \goCode\src\排序\main.go
 * @Description: 实现快排，归并排序，堆排序，二分搜索
 */

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	array := []int{1, 4, 6, 1, 2, 88, 22, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 12, 352, 321, 674, 2, 5, 123, 125}
	fmt.Println(quickSort(array), len(array))
	//fmt.Println(mergeSort(array))
	//fmt.Println(heapSort(array))
	a := quickSort(array)
	index := binarySearch(a, 321)
	fmt.Println(index)
}

/**
 * @description: 快排
 * @param {[]int} arr
 * @return {*}
 */
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := []int{}
	right := []int{}

	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[0] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(quickSort(left), append([]int{arr[0]}, quickSort(right)...)...)
}

/**
 * @description: 归并排序
 * @param {[]int} arr
 * @return {*}
 */
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else if len(arr) == 2 {
		if arr[0] > arr[1] {
			return []int{arr[1], arr[0]}
		}
		return arr
	}

	l := 0
	r := len(arr) - 1
	mid := (l + r) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	result := []int{}
	i, j := 0, 0
	for {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else if left[i] > right[j] {
			result = append(result, right[j])
			j++
		}
		if i == len(left) {
			result = append(result, right[j:]...)
			break
		}
		if j == len(right) {
			result = append(result, left[i:]...)
			break
		}
	}
	return result
}

/**
 * @description: 小根堆排序
 * @param {[]int} arr
 * @return {*}
 */
func heapSort(arr []int) []int {
	h := &Myheap{}
	heap.Init(h)
	result := []int{}
	for i := range arr {
		heap.Push(h, arr[i])
	}
	for len(*h) > 0 {
		result = append(result, heap.Pop(h).(int))
	}
	return result
}

type Myheap []int

func (h Myheap) Len() int           { return len(h) }
func (h Myheap) Less(i, j int) bool { return h[i] < h[j] }
func (h *Myheap) Swap(i, j int)     { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *Myheap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
func (h *Myheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *Myheap) Top() int {
	return (*h)[0]
}

/**
 * @description: 二分查找
 * @param {[]int} arr
 * @param {int} target
 * @return {*}
 */
func binarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := (l + r) >> 1
		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
