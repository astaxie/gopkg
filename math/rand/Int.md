## func Int() int

返回值：

- 返回int

功能说明：


这个函数主要实现返回一个非负伪随机整数.


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
			x := rand.Int()
			fmt.Println(x)
			i += 1
		}
	}







