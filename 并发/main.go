/*
 * @Author: Alexleslie
 * @Date: 2022-03-06 22:49:53
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-07 01:31:44
 * @FilePath: \goCode\src\并发\main.go
 * @Description: 实现简单的并发编程
 */
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

func main() {
	Simple()
	withContext()
}

func Simple() {
	var wg sync.WaitGroup
	wg.Add(3)
	ch1 := make(chan interface{}, 1)
	ch2 := make(chan interface{}, 1)
	ch3 := make(chan interface{}, 1)

	iter := 1000
	fruits := []string{"apple", "pear", "watermelon"}
	ch1 <- []int{}

	go genFruits(fruits[0], ch1, ch2, iter, &wg)
	go genFruits(fruits[1], ch2, ch3, iter, &wg)
	go genFruits(fruits[2], ch3, ch1, iter, &wg)

	wg.Wait()
}

/**
 * @description: 使用空结构体进行参数传递，省空间，并不关闭channel，等gc回收
 * @param {string} name
 * @param {*} ch1
 * @param {chaninterface{}} ch2
 * @param {int} iter
 * @param {*sync.WaitGroup} wg
 * @return {*}
 */
func genFruits(name string, ch1, ch2 chan interface{}, iter int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < iter; i++ {
		select {
		case <-ch1:
			fmt.Printf("%d: %v \n", i+1, name)
			ch2 <- []int{}
		}
	}
}

/**
 * @description: 使用context控制子线程
 * @param {context.Context} ctx
 * @param {context.CancelFunc} cancel
 * @param {string} name
 * @param {*} ch1
 * @param {chaninterface{}} ch2
 * @param {int} iter
 * @param {*sync.WaitGroup} wg
 * @return {*}
 */
func genFruitsCtx(ctx context.Context, cancel context.CancelFunc, name string, ch1, ch2 chan interface{}, iter int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		select {
		case <-ch1:
			fmt.Printf("%d %v \n", i, name)
			i += 1
			if i == iter && name == "watermelon" {
				cancel()
			} else {
				ch2 <- []int{}
			}
		case <-ctx.Done():
			close(ch1) // 关闭channel，但是不关闭也行，gc会自动回收
			log.Printf("%v 's go routine is end \n", name)
			return
		}
	}

}

func withContext() {
	var wg sync.WaitGroup
	wg.Add(3)
	ch1 := make(chan interface{}, 1)
	ch2 := make(chan interface{}, 1)
	ch3 := make(chan interface{}, 1)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	iter := 1000
	fruits := []string{"apple", "pear", "watermelon"}
	ch1 <- []int{}

	go genFruitsCtx(ctx, cancel, fruits[0], ch1, ch2, iter, &wg)
	go genFruitsCtx(ctx, cancel, fruits[1], ch2, ch3, iter, &wg)
	go genFruitsCtx(ctx, cancel, fruits[2], ch3, ch1, iter, &wg)

	wg.Wait()

}
