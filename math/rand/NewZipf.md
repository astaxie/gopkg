## func NewZipf(r *Rand, s float64, v float64, imax uint64) *Zipf

参数列表:

- r *Rand 随机值产生器结构体指针
- s float64 Zipf公式的描述分布的指数
- v float64 Zipf公式的v参数
- imax uint64 分布区间的最大值


返回值:

- Zipf对象的指针

功能说明：

该函数实现取值是0到imax之间的随即变量p(k), 它与(v+k)**(-s)成比例, 其中s>1, k>=0, 且 v>=1.

![](http://upload.wikimedia.org/math/f/7/e/f7ea651d890ccb3249ecd1c8eded8869.png)

![](http://upload.wikimedia.org/math/3/4/4/344c5375a5cd6ea108c6823b7b167b02.png)

更多关于Zipf分布的理论请参考[wikipedia](http://en.wikipedia.org/wiki/Zipf's_law)


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







