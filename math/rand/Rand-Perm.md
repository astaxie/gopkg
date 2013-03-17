## func (*Rand) Perm(n int) []int

参数列表

- n 期望输出随机值的最大限制值

返回值：

- 返回[]int

功能说明：

返回一个元素个数是n的slice,里面的元素是0到n-1的整数无重复的随机排列.

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
			x := r.Perm(10)
			fmt.Println(x)
			i += 1
		}
	}









