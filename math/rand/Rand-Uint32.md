## func (*Rand) Uint32() uint32

返回值：

- 返回int32

功能说明：


这个函数主要实现返回一个0-2^32范围内的伪随机数.


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
			x := r.Uint32()
			fmt.Println(x)
			i += 1
		}
	}








