/*
 * @Author: Alexleslie
 * @Date: 2022-03-09 02:14:25
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-09 02:31:36
 * @FilePath: \goCode\src\随机数生成\main.go
 * @Description: 随机数生成，利用rand7 实现 rand10。
 * 利用rand7实现的关键是生成10个概率相同的数 = 生成十个数的次数相同
 * 我们以十进制为例子，先随机生成个位数，再随机生成十位数，两个数互不干扰，所以保证生成的数次数都是一
 * 个位数0~9，十位数0~9 取一个数，其个位数的概率为1/10，十位数的概率为1/10，所以实现了概率相同
 * 调用一次生成7个数，两次生成7*7个数


 */

package main

import "math/rand"

func main() {

}

func rand10() int {
	for {
		a := rand7() - 1
		b := (rand7()-1)*7 + a // 总共产生了49个数，包含0
		if b < 40 {            // 取40个数，需要是10的倍数
			return b%10 + 1
		}
	}
}

func rand7() int {
	return rand.Intn(6) + 1
}
