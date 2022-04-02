/*
 * @Author: Alexleslie
 * @Date: 2022-03-17 19:09:00
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-31 18:16:29
 * @FilePath: \src\测试\main.go
 * @Description:
 */

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var base float64 = 10000000
	rate := 0.0001 // 缺失率
	data := make([]int, int(base))

	// 理解为base个块数的文件，并随机抽取1%作为坏块, 坏块标记1
	for i := 0; i < int(float64(base)*rate); i++ {
		po := random(int(base))
		data[po] = 1
	}
	total := 10000
	count := 4600
	catch := 0
	// 进行total次实验，每次审计count个块数
	for j := 0; j < total; j++ {
		// 进行count 块数审计，如果命中则catch + 1
		for i := 0; i < count; i++ {
			po := random(int(base))
			if data[po] == 1 {
				catch++
				break
			}
		}
	}
	//fmt.Printf("Background Situation: totalBlock: %d, totalSize: %v MB, errorRate: %v , errorSize: %v MB \n", int(base), toMb(base), rate, toMb(base*rate))
	//fmt.Printf("ChallengeBlock :%v, transferSize: %v MB, CorrectRate: %v \n", count, toMb(float64(count)), float64(catch)/float64(total))
	// getP(base, float64(count), 100, 1)
	// getP(base, float64(count), 100, 2)
	// getP(base, float64(count), 100, 3)
	// getP(base, float64(count), 10000, 1)
	// getP(base, float64(count), 10000, 2)
	// getP(base, float64(count), 10000, 3)
	// getP(base, float64(count), 10, 3)
	ans := getTransfer(base, float64(count), 10000, count)
	fmt.Println(ans / 1024)

}

func getP(base, c float64, n, k int) float64 {
	p := c / base
	t := p
	p = 1
	for i := 1; i <= k; i++ {
		p = p * t
	}
	for i := k; i < n; i++ {
		p = p * (1 - t)
	}
	a := 1
	temp := 0
	for temp < k {
		a = a * (n - temp)
		temp++
	}
	b := 1
	temp = 0
	for temp < k {
		b = b * (k - temp)
		temp++
	}
	combine := a / b
	//fmt.Printf("n: %d, k: %d p:%v \n", n, k, float64(combine)*p)
	return float64(combine) * p
}

func getTransfer(base, c float64, n, count int) float64 {
	var ans float64
	temp := 1 - (getP(base, c, n, 0))
	for i := 0; i < count; i++ {
		ans += getP(base, c, n, 1)/temp*32 + getP(base, c, n, 2)/temp*32 + getP(base, c, n, 3)/temp*32 + getP(base, c, n, 4)/temp*4096
	}
	return ans
}

func random(base int) int {
	max := big.NewInt(int64(base))
	po, _ := rand.Int(rand.Reader, max)
	return int(po.Int64())
}

func toMb(size float64) float64 {
	return size * (1 << 12) / (1 << 20)
}
