## func IsSorted(data Interface) bool

参数列表

- data 表示要判断的 Interface 数据

返回值：

- 返回 bool

功能说明：

		IsSorted 判断数据是否已经按升序排列

代码案例：

	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		d := []int{5, 2, 6, 3, 1, 4}                 // unsorted
		fmt.Println(sort.IsSorted(sort.IntSlice(d))) // false
	
		sort.Sort(sort.IntSlice(d))
		fmt.Println(sort.IsSorted(sort.IntSlice(d))) // true
	
		a := []float64{5.5, 2.2, 6.6, 3.3, 1.1, 4.4}
		fmt.Println(sort.IsSorted(sort.Float64Slice(a))) // false
	
		sort.Sort(sort.Float64Slice(a))
		fmt.Println(sort.IsSorted(sort.Float64Slice(a))) // true
	
		s := []string{"PHP", "golang", "java", "python", "C", "Objective-C"}
		fmt.Println(sort.IsSorted(sort.StringSlice(s))) // false
	
		sort.Sort(sort.StringSlice(s))
		fmt.Println(sort.IsSorted(sort.StringSlice(s))) // true
	}
	


