/*
 * @Author: Alexleslie
 * @Date: 2022-03-28 01:15:25
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-28 01:28:19
 * @FilePath: \src\pattern\singleton.go
 * @Description:
 */

package pattern

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Singleton interface {
	SaySomething()
}

type singleton struct {
	Val1 int
	Val2 int
}

func (singleton) SaySomething() {
	fmt.Println("Singleton")
}

var singletonInstance = &singleton{random(10000), random(10000)}

func NewSingletonInstance() Singleton {
	return singletonInstance
}

func random(base int) int {
	max := big.NewInt(int64(base))
	po, _ := rand.Int(rand.Reader, max)
	return int(po.Int64())
}
