## func Float64() float64

返回值：

- 返回float64

功能说明：


这个函数主要实现返回一个[0.0.1.0)区间的Float64型的数.通常是一个精确到小数点后16位(16x4=64)的浮点数.


代码实例：

	package main

	import (
		"fmt"
		"math/rand"
	)

	func main() {
		n := 10
		i := 0
		for i < n {
			x := rand.Float64()
			fmt.Println(x)
			i += 1
		}
	}






