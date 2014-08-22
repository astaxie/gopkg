##	func Reverse(data Interface) Interface

参数列表

- data 表示要逆序的 Interface 数据

返回值：

- 返回 Interface

功能说明：

- Reverse 返回逆序的Inferface数据

代码案例： 
		
	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		s := []int{5, 2, 6, 3, 1, 4} // unsorted
		sort.Sort(sort.Reverse(sort.IntSlice(s)))
		fmt.Println(s) // [6 5 4 3 2 1]
	}
	
	
												