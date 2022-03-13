/*
 * @Author: Alexleslie
 * @Date: 2022-03-13 11:51:09
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-13 13:56:07
 * @FilePath: \goCode\src\数据结构\Heap.go
 * @Description: 实现大顶堆与小顶堆
 */

package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 12345, 1, 35, 132, 6, 4, 13, 12, 3, 4, 34, 1323, 5126, 3, 41, 3, 5, 126, 58, 79783, 5413, 825, 746, 4358, 5687, 3, 6}
	fmt.Println(heapSort(list))
}

func buildHeap(l []int) []int {
	for i := len(l)/2 - 1; i >= 0; i-- {
		heapify(l, i)
	}
	return l
}

func heapify(l []int, start int) { // 主要是个递归的操作
	if start >= len(l) {
		return
	}
	c1 := start*2 + 1
	c2 := start*2 + 2
	max := start
	if c1 < len(l) && l[c1] > l[max] {
		max = c1
	}
	if c2 < len(l) && l[c2] > l[max] {
		max = c2
	}
	if max != start {
		l[start], l[max] = l[max], l[start]
		heapify(l, max)
	}
}

func Pop(l *[]int) int {
	(*l)[0], (*l)[len(*l)-1] = (*l)[len(*l)-1], (*l)[0]
	max := (*l)[len(*l)-1]

	*l = (*l)[:len(*l)-1]
	heapify(*l, 0)
	return max
}

func heapSort(l []int) []int {
	result := []int{}
	hp := buildHeap(l)
	for len(hp) > 0 {
		x := Pop(&hp)
		result = append(result, x)
	}
	return result
}
