## func (*Rand) NormFloat64() float64

返回值：

- 返回float64

功能说明：

该函数返回一个区间为[-math.MaxFloat64, +math.MaxFloat64]的标准正态分布.

正太分布的密度函数是

![](http://upload.wikimedia.org/math/7/9/a/79af499be9466b7dce2cf8ac19fa0a07.png)

其中默认的位置参数(期望desiredMean)为0,尺度参数(标准差desiredStdDev)为1.
如果想改变不同的位置参数和尺度参数,可以通过下面方法变通:

	sample = NormFloat64() * desiredStdDev + desiredMean
	

代码实例：

	package main

	import (
		"fmt"
		"math/rand"
	)

	func main() {
		n := 10000
		i := 0
		sum := 0.0
			r := rand.New(rand.NewSource(64))
		for i < n {
			x := r.NormFloat64()
			i += 1
			sum += x
		}
		expect := sum / float64(n)
		fmt.Println(expect)
	}










