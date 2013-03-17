## func NewZipf(r *Rand, s float64, v float64, imax uint64) *Zipf

参数列表:

- r *Rand 随机值产生器结构体指针
- s float64 Zipf公式的s参数
- v float64 Zipf公司的v参数
- imax uint64 分布区间的最大值


返回值:

- Zipf 指针

功能说明：

该函数实现取值是0到imax之间与(v+k)**(-s)成比例的Zipf分布.


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







