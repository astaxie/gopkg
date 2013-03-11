## func Sort(data Interface)

参数列表

- data 表示要排序的 Interface 数据

功能说明：

Sort 对 data 进行排序。 它调用一次 data.Len 来决定排序的长度 n，调用 data.Less 和 data.Swap 的开销为 O(n*log(n))。此排序为不稳定排序。

代码案例：

	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		d := []int{5, 2, 6, 3, 1, 4} // unsorted
		sort.Sort(sort.IntSlice(d))
		fmt.Println(d)
		// Output:[1 2 3 4 5 6]
	
		a := []float64{5.5, 2.2, 6.6, 3.3, 1.1, 4.4}
		sort.Sort(sort.Float64Slice(a))
		fmt.Println(a)
		// Output:[1.1 2.2 3.3 4.4 5.5 6.6]
		
		s := []string{"PHP", "golang", "java", "python", "C", "Objective-C"}
		sort.Sort(sort.StringSlice(s))
		fmt.Println(s)
		// Output:[C Objective-C PHP golang java python]
	}
	