## func IntsAreSorted(a []int) bool

参数列表

- a 表示需要判断的 int 切片

返回值：

- 返回bool

功能说明：

IntsAreSorted 判断 int 切片是否已经按升序排列。

代码实例：

	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		s := []int{5, 2, 6, 3, 1, 4}       // unsorted
		fmt.Println(sort.IntsAreSorted(s)) // false
		sort.Ints(s)
		fmt.Println(sort.IntsAreSorted(s)) // true
	}