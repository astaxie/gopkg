## func Ints(a []int)

参数列表

- a 表示要排序的 int 切片

功能说明：

Ints 以升序排列 int 切片

代码实例：

	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		s := []int{5, 2, 6, 3, 1, 4} // unsorted
		sort.Ints(s)
		fmt.Println(s) // [1 2 3 4 5 6]
	}
	

