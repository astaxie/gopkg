## func New(src Source) *Rand

参数列表

- src Source

返回值

- Rand指针

该函数主要返回了一个新的Rand实例,并以src作为随机值产生器.


代码实例

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
			fmt.Println(r.Int())
			i += 1
		}
	}






