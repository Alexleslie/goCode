/*
 * @Author: Alexleslie
 * @Date: 2022-03-28 01:15:52
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-28 01:25:13
 * @FilePath: \src\main.go
 * @Description:
 */
package main

import (
	"fmt"
	"pattern"
)

func main() {
	ch := make(chan int, 2)
	go func() {
		instance := pattern.NewSingletonInstance()
		fmt.Println(instance)
		ch <- 1
	}()
	go func() {
		instance2 := pattern.NewSingletonInstance()
		fmt.Println(instance2)
		ch <- 1
	}()
	<-ch
	<-ch
}
