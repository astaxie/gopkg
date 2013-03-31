## func (*Rand) Float32() float32

返回值：

- 返回float32

功能说明：


这个函数主要实现返回一个[0.0.1.0)区间的Float32型的数.通常是一个精确到小数点后8位(8x4=32)的浮点数.


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
			x := r.Float32()
			fmt.Println(x)
			i += 1
		}
	}







