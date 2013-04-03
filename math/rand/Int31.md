## func Int31() int32

返回值：

- 返回int32

功能说明：


这个函数主要实现返回一个非负伪随机整数,该整数是一个有符号int32型,除去最高位,占31个比特位,范围在[0,2^31-1]之间.


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
			x := rand.Int31()
			fmt.Println(x)
			i += 1
		}
	}







