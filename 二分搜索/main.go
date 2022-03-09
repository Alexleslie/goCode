/*
 * @Author: Alexleslie
 * @Date: 2022-03-09 14:19:06
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-09 16:18:08
 * @FilePath: \goCode\src\二分搜索\main.go
 * @Description: 二分搜索的用例
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{1, 4, 6, 1, 2, 88, 22, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 12, 352, 321, 674, 2, 5, 123, 125}
	array2 := []int{5, 3, 4, 7, 1, 3, 5, 67, 34, 54, 12, 45, 15, 1245, 45, 45, 1, 15, 15, 45, 45, 15, 46, 7, 8, 769, 659, 50, 58, 346, 27, 8, 2}
	sort.Ints(array)
	sort.Ints(array2)
	temp := make([]int, len(array)+len(array2))
	copy(temp[:len(array)], array)
	copy(temp[len(array):], array2)
	sort.Ints(temp)

	k := 60
	fmt.Println(temp[k-1])
	fmt.Println(findKthlargest(array, array2, k))
}

func findKthlargest(arr1, arr2 []int, k int) int {
	l1, l2 := 0, 0
	for {
		if l1 == len(arr1) {
			return arr2[l2+k-1]
		}
		if l2 == len(arr2) {
			return arr1[l1+k-1]
		}
		if k == 1 {
			return min(arr1[l1], arr2[l2])
		}
		half := k / 2
		newl1 := min(l1+half, len(arr1)) - 1 // 每次移动的距离是half，但是遇到边界需要停下
		newl2 := min(l2+half, len(arr2)) - 1

		num1, num2 := arr1[newl1], arr2[newl2]
		if num1 <= num2 {
			k -= (newl1 - l1 + 1) // 真实移动距离
			l1 = newl1 + 1
		} else {
			k -= (newl2 - l2 + 1)
			l2 = newl2 + 1
		}
	}

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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
