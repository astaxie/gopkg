## func (*Rand) Seed(seed int64)

参数列表

- seed 提供随机种子值

功能说明：

以提供的参数作为种子值,来初始化随机产生器.如果Seed没有被调用,那么随机数调用前默认调用Seed(1).

小贴士,一般用传入time.Now().UnixNano()给Seed函数来实现不同运行次数看到不同结果的目的.

代码实例：

	package main

	import (
		"fmt"
		"math/rand"
		"time"
	)

	func main() {
		n := 10
		i := 0
		r := rand.New(rand.NewSource(64))
		for i < n {
			r.Seed(time.Now().UnixNano())
			fmt.Println(r.Int())
			i += 1
		}
	}









