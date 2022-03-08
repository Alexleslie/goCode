/*
 * @Author: Alexleslie
 * @Date: 2022-03-08 19:45:51
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-08 20:17:37
 * @FilePath: \goCode\src\全排列与组合\main.go
 * @Description:
 */

package main

import "fmt"

func main() {
	str := "abc"
	fmt.Println(allPermutate(str))

}

func allPermutate(str string) []string {
	result := []string{}
	dict := make(map[byte]bool)

	var permu func(temp string)
	permu = func(temp string) {
		if len(temp) == len(str) {
			result = append(result, temp)
			return
		}
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
