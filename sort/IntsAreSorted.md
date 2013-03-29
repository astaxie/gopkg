## func IntsAreSorted(a []int) bool

参数列表

- a 表示要判断的 int 切片

返回值：

- 返回 bool

功能说明：

IntsAreSorted 判断 int 切片是否已经按升序排列。

代码实例：

	package main
	
	import (
	"fmt"
	"sort"
	)
	
	func main() {
		a := []int{5, 2, 6, 3, 1, 4}       // unsorted
		fmt.Println(sort.IntsAreSorted(a)) // false
		sort.Ints(a)
		fmt.Println(a)                     // [1 2 3 4 5 6]
		fmt.Println(sort.IntsAreSorted(a)) // true
	}
	
