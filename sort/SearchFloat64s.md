## func SearchFloat64s(a []float64, x float64) int

参数列表

- a 表示升序排列的 float64 切片
- x 表示搜索的 float64 值

返回值：

- 返回索引

功能说明：

SearchFloat64s 在 float64s 切片中搜索 x 并返回索引，如 Search 函数所述. 返回可以插入 x 值的索引位置，如果 x 不存在，返回数组a的长度。
切片必须以升序排列。

代码案例：

	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		a := [10]float64{-29.9, -19.9, -9.9, 0, 9, 19.9, 29, 39.9, 	49.9, 59}
	
		fmt.Println(sort.SearchFloat64s(a[:], -15)) // 2
		fmt.Println(sort.SearchFloat64s(a[:], 0.9)) // 4
		fmt.Println(sort.SearchFloat64s(a[:], 25))  // 6
		fmt.Println(sort.SearchFloat64s(a[:], 49))  // 8
	}	
	
