## 结构 StringSlice 

====

		type StringSlice []String
		
StringSlice 针对 []String 实现接口的方法，以升序排列
	
StringSlice 有以下方法：

====
- func (p StringSlice) Len() int

	返回 p 的长度
	
====
- func (p StringSlice) Less(i, j int) bool 

	返回 p[i] < p[j] 是否为真
	
====
- func (p StringSlice) Search(x String) int

	返回小于 x 的最小元素的索引

====
- func (p StringSlice) Swap(i, j int)
	
	交换 p[i]和p[j] 的值

====
- func (p StringSlice) Sort() 

	按升序排列 p 切片
	
代码案例：
	
	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		p := sort.StringSlice{"PHP", "golang", "java", "python", "C", "Objective-C"}
	
		fmt.Println(p.Len()) // 6
	
		fmt.Println(p.Less(0, 1)) // true
	
		fmt.Println(p.Search("go")) // 1
	
		p.Swap(0, 1)
		fmt.Println(p)
		// Output: [golang PHP java python C Objective-C]
	
		p.Sort()
		fmt.Println(p)
		// Output: [C Objective-C PHP golang java python]
	}
	
