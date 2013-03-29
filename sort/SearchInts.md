## func SearchInts(a []int, x int) int

参数列表

返回值：

- a 表示升序排列的 int 切片 
- x 表示搜索的 int 值

返回值：

- 返回索引

功能说明：

SearchInts 在 ints 切片中搜索x并返回索引，如 Search 函数所述。返回可以插入 x 值的索引位置，如果 x 不存在，返回数组 a 的长度。切片必须以升序排列。

代码案例：
	
	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		a := [10]int{-29, -19, -9, 0, 9, 19, 29, 39, 49, 59}
		
		fmt.Println(sort.SearchInts(a[:], -15)) // 2
		fmt.Println(sort.SearchInts(a[:], 5))   // 4
		fmt.Println(sort.SearchInts(a[:], 25))  // 6
		fmt.Println(sort.SearchInts(a[:], 49))  // 8
	} 
		