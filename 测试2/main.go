/*
 * @Author: Alexleslie
 * @Date: 2022-04-01 14:05:03
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-04-01 14:06:22
 * @FilePath: \src\测试2\main.go
 * @Description:
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()

	time.Sleep(time.Second)

	spend := time.Now().Sub(a)
	fmt.Println(spend)
	fmt.Println(spend.Seconds())

}
