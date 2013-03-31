## func Int63n(n int64) int64

参数列表

- n 期望输出随机值的最大限制值

返回值：

- 返回int64

功能说明：


返回一个[0,n)之间的非负伪随机整数.如果n<=0会panic.

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
			x := rand.Int63n(5)
			fmt.Println(x)
			i += 1
		}
	}








