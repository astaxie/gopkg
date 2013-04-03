## func NewSource(seed int64) Source

参数列表:

- int64类型的值作为种子值


返回值:

- Source结构体

功能说明：
	
该函数主要返回一个指定种子的随机数产生器.


代码实例:

	package main

	import (
		"fmt"
		"math/rand"
		"time"
	)

	func main() {
		n := 10
		i := 0

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i < n {
			fmt.Println(r.Int())
			i += 1
		}
	}







