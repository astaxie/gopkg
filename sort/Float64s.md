## func Float64s(a []float64)

参数列表

- a 表示未排序的float64切片

功能说明：

Float64s 以升序排列float64切片

代码实例：

	package main
	
	import (
		"fmt"
		"sort"
	)
		
	func main() {
		a := []float64{5.5, 2.2, 6.6, 3.3, 1.1, 4.4} // unsorted
		sort.Float64s(a)
		fmt.Println(a)
		// Output: [1.1 2.2 3.3 4.4 5.5 6.6]
	}




