## func (*Rand) Int31n(n int32) int32

参数列表

- n 期望输出随机值的最大限制值

返回值：

- 返回int32

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
		r := rand.New(rand.NewSource(64))
		for i < n {
			x := r.Int31n(5)
			fmt.Println(x)
			i += 1
		}
	}









