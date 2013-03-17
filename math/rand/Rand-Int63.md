## func (*Rand) Int63() int64

返回值：

- 返回int64

功能说明：


这个函数主要实现返回一个非负伪随机整数,该整数是一个有符号int64型,除去最高位,占63个比特位,范围在[0,2^63-1]之间.


代码实例：

	package main

	import (
		"fmt"
		"math/rand"
	)

	func main() {
		n := 10
		i := 0
			r := rand.New(rand.NewSource(64))
		for i < n {
			x := r.Int63()
			fmt.Println(x)
			i += 1
		}
	}








