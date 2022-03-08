/*
 * @Author: Alexleslie
 * @Date: 2022-03-08 22:43:10
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-08 22:50:23
 * @FilePath: \goCode\src\滑动窗口\main.go
 * @Description:
 */

package main

import "fmt"

func main() {
	str := "abcabcbb"
	fmt.Println(longestSubstring(str))
}

/**
 * @description: 使用滑动窗口实现最长无重复子串，滑动窗口最重要的是左指针和右指针，也就是左边界和右边界
 * @param {string} s
 * @return {*}
 */
func longestSubstring(s string) int {
	dict := make(map[byte]bool)
	n := len(s)
	l := 0
	r := 0
	max := 0
	for r < n {
		// 当窗口右边界在范围里面时
		word := s[r]
		if dict[word] == false {
			dict[word] = true
			r++
			if r-l > max { // 计算每一次移动后窗口的大小
				max = r - l
			}
		} else {
			// 发现重复后，左边界移到重复字符的右边一格
			// 如 a(l)bcdbe(r) -> abc(l)dbe(r)
			for i := l; i <= r; i++ {
				if s[i] != word {
					dict[s[i]] = false
				} else {
					l = i + 1
					r++
					break
				}
			}
		}
	}
	return max
}
