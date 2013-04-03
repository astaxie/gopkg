## func (*Zipf) Uint64 uint64


返回值:

- unit64

功能说明：

该函数返回一个按照Zipf对象表述的Zipf分布的值.


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
		zipf := rand.NewZipf(r, 3.14, 2.72, 5000)
		for i < n {
			fmt.Println(zipf.Uint64())
			i += 1
		}
	}







