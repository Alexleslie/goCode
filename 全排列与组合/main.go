/*
 * @Author: Alexleslie
 * @Date: 2022-03-08 19:45:51
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-08 20:59:28
 * @FilePath: \goCode\src\全排列与组合\main.go
 * @Description: 实现全排列与任意长度组合
 */

package main

import (
	"fmt"
)

func main() {
	str := "abc"
	fmt.Println(allPermutate(str))
	fmt.Println(Combine(str))

}

/**
 * @description: 实现全排列
 * @param {string} str
 * @return {*}
 */
func allPermutate(str string) []string {
	result := []string{}
	dict := make(map[byte]bool)

	var permu func(temp string)
	permu = func(temp string) {
		if len(temp) == len(str) {
			result = append(result, temp)
			return
		}

		// a->b->c; a->c->b ..
		for i := 0; i < len(str); i++ {
			cur := str[i]
			if dict[cur] == false {
				temp += string(cur)
				dict[cur] = true
				permu(temp)
				dict[cur] = false
				temp = temp[:len(temp)-1]
			}
		}
	}
	permu("")
	return result
}

/**
 * @description: 生成所有可能长度的组合，与全排列十分类似，只需要加上固定前后顺序，所以只需要设定遍历的起始顺序
 * @param {string} str
 * @return {*}
 */
func Combine(str string) []string {
	result := []string{}
	dict := make(map[byte]bool)

	var permu func(temp string, size, index int)

	permu = func(temp string, size, index int) {
		if len(temp) == size {
			result = append(result, temp)
		}

		// a,b,c; a->b, a->c, b->c
		for i := index; i < len(str); i++ {
			cur := str[i]
			if dict[cur] == false {
				temp += string(cur)
				dict[cur] = true
				permu(temp, size, i+1)
				// start from i+1
				temp = temp[:len(temp)-1]
				dict[cur] = false
			}
		}
	}
	for i := 1; i <= len(str); i++ {
		permu("", i, 0)
	}
	return result
}
