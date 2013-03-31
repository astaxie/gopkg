## func ExpFloat64() float64

返回值：

- 返回float64

功能说明：


这个函数主要实现了指数分布公式

![](http://upload.wikimedia.org/math/7/c/1/7c1e7458e99f77f22c350aec59c67e9c.png)

函数定义域在(0, +∞);在此,用math.MaxFloat64表示正无穷, 率参数λ归元为1,函数的期望是1/λ(λ=1).返回值计算如下:

	f(x,1)


如果想要λ是别的值怎么办呢?可以通过调整分母来达到目的.


	sample = ExpFloat64() / desiredRateParameter


代码实例：

	package main

	import (
		"fmt"
		"math/rand"
	)

	func main() {
		sum := 0.0
		n := 10000000
		i := 0
		for i < n {
			x := rand.ExpFloat64()
			sum += x
			i += 1
		}
		expect := sum / (float64)(n)
		fmt.Println(expect)
	}




