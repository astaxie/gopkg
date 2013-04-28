## 结构 IntSlice

====

		type IntSlice []int
		
IntSlice 针对 []int 实现接口的方法，以升序排列
	
IntSlice 有以下方法：

====
- func (p IntSlice) Len() int

	返回 p 的长度
	
====
- func (p IntSlice) Less(i, j int) bool 

	返回 p[i] < p[j] 是否为真
	
====
- func (p IntSlice) Search(x int) int

	返回小于 x 的最小元素的索引

====
- func (p IntSlice) Swap(i, j int)
	
	交换 p[i]和p[j] 的值

====
- func (p IntSlice) Sort() 

	按升序排列 p 切片
	
代码案例：
	
	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		a := sort.IntSlice{5, 2, 6, 3, 1, 4} // unsorted
		
		fmt.Println(a.Len()) // 6
		
		fmt.Println(a.Less(0, 1)) // false
		
		fmt.Println(a.Search(4)) // 5
		
		a.Swap(0, 1)
		fmt.Println(a) // Output: [2 5 6 3 1 4]
		
		a.Sort()
		fmt.Println(a) // Output: [1 2 3 4 5 6]	
	}
	