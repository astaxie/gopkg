## func Float64sAreSorted(a []float64) bool

参数列表

- a 表示需要判断的 float64 切片

返回值：

- 返回bool

功能说明：

Float64sAreSorted 判断 float64 切片是否已经按升序排列。

代码实例：

	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		a := []float64{5.5, 2.2, 6.6, 3.3, 1.1, 4.4} 
		fmt.Println(sort.Float64sAreSorted(a)) // false
		sort.Float64s(a)
		fmt.Println(sort.Float64sAreSorted(a)) // true
	}
