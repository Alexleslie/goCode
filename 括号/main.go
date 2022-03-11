/*
 * @Author: Alexleslie
 * @Date: 2022-03-11 12:05:08
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-11 12:19:35
 * @FilePath: \goCode\src\括号\main.go
 * @Description: 实现解决括号匹配的各种情况，括号匹配首先想用使用栈
 */

package main

import (
	"fmt"
)

func main() {
	str := ")())(())))(()())(())((()()("
	fmt.Println(longestValidParentheses(str))
	fmt.Println(longestValidParentheses2(str))

}

func longestValidParentheses2(s string) int {
	maxAns := 0
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns

}

/**
 * @description: 动态规划实现，分两种情况，(), ((...))  我们只需要关注右括号。
 * @param {string} s
 * @return {*}
 */
func longestValidParentheses(s string) int {
	maxAns := 0
	dp := make([]int, len(s)+1)
	dp[0] = 0 // 包含该括号的最大有效括号长度
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i+1] = dp[i-1] + 2
			} else {
				if i-dp[i]-1 < 0 {
					dp[i] = 0
				} else if s[i-dp[i]-1] == '(' {
					dp[i+1] = dp[i] + dp[i-dp[i]-1] + 2
				} else {
					dp[i+1] = 0
				}
			}
		} else {
			dp[i+1] = 0
		}
		maxAns = max(maxAns, dp[i+1])
	}
	fmt.Println(dp)
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
