## 结构 Float64Slice

====
		type Float64Slice []float64
		
Float64Slice 针对 []Float64Slice 实现接口的方法，以升序排列
	
Float64Slice 有以下方法：

====
- func (p Float64Slice) Len() int

	返回 p 的长度
	
====
- func (p Float64Slice) Less(i, j int) bool 

	返回 p[i] < p[j] 是否为真
	
====
- func (p Float64Slice) Search(x float64) int

	返回小于 x 的最小元素的索引

====
- func (p Float64Slice) Swap(i, j int)
	
	交换 p[i]和p[j] 的值

====
- func (p Float64Slice) Sort() 

	按升序排列 p 切片
	
代码案例：
	
	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		a := sort.Float64Slice{5.5, 2.2, 6.6, 3.3, 1.1, 4.4} // unsorted
	
		fmt.Println(a.Len()) // 6
	
		fmt.Println(a.Less(0, 1)) // false
		
		fmt.Println(a.Search(4.0)) // 5
		
		a.Swap(0, 1)
		fmt.Println(a)
		// Output: [2.2 5.5 6.6 3.3 1.1 4.4]
	
		a.Sort()
		fmt.Println(a)
		// Output: [1.1 2.2 3.3 4.4 5.5 6.6]
	}
	





