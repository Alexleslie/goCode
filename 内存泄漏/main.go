/*
 * @Author: Alexleslie
 * @Date: 2022-03-07 01:32:35
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-07 17:50:45
 * @FilePath: \goCode\src\内存泄漏\main.go
 * @Description: golang中内存泄漏的例子：内存泄漏指的是goroutine使用的内存没有得到回收与释放，
 	使得内存占用一直上升。
	内存泄漏的排查方法，通过在两个时间点之间查看heap的堆栈情况
	时间点1：go tool pprof http://localhost:6060/debug/pprof/heap
	时间点2：go tool pprof http://localhost:6060/debug/pprof/heap
	排查关于单个函数所占用的空间
*/

package main

import (
	"fmt"
	"sync"
	"time"

	// "github.com/pkg/profile"
	"net/http"
	_ "net/http/pprof"
)

var global [10000]string // 定义一个全局变量
var mutex sync.Mutex

func main() {
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop() // 看内存泄漏用处不大
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		http.ListenAndServe("localhost:6060", nil)
		wg.Done()

	}()
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second)
		go func(k int) { // 闭包中的变量都是传地址的 闭包是函数和它所引用的环境
			defer wg.Done()
			s := str()
			// 在goroutine里生成字符串，此时字符串还是存储在栈内

			//temp := s[:2]
			// temp和s的底层数组是同一个

			temp := string([]byte(s[:2]))
			// 通过强制类型转换来实现对s底层数组的复制，s此后可以被回收

			mutex.Lock()
			global[k] = temp
			mutex.Unlock()
			// 往全局变量添加temp后，原先的s后面的字符串部分并没有的得到释放

		}(i)
	}
	wg.Wait()
	fmt.Println(global, len(global))
}

func str() string {
	r := ""
	for i := 0; i < 10000; i++ {
		r += "hello"
	}
	return r

}
