## func Sort(data Interface)

参数列表

- data 表示要排序的 Interface 数据

功能说明：

Sort 对 data 进行排序。 它调用一次 data.Len 来决定排序的长度 n，调用 data.Less 和 data.Swap 的开销为 O(n*log(n))。此排序为不稳定排序。

====
	 type Interface interface {	
 	   // Len is the number of elements in the collection.
  	  // Len 为集合内元素的总数
 	   Len() int
 	   // Less returns whether the element with index i should sort
 	   // before the element with index j.
  	  // Less 返回索引为 i 的元素是否应排在索引为 j 的元素之前。
  	  Less(i, j int) bool
  	  // Swap swaps the elements with indexes i and j.
  	  // Swap 交换索引为 i 和 j 的元素
  	  Swap(i, j int)
 	}
		
任何实现了 sort.Interface 的类型（一般为集合），均可使用该包中的方法进行排序。这些方法要求集合内列出元素的索引为整数。

代码案例（一）：

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
	

代码案例（二）：

	package main
	
	import (
	"fmt"
	"sort"
	)
	
	type ByLength []string
	
	func (s ByLength) Len() int {
		return len(s)
	}
	
	func (s ByLength) Swap(i, j int) {
		s[i], s[j] = s[j], s[i]
	}
	
	func (s ByLength) Less(i, j int) bool {
		return s[i] < s[j]
	}
	
	func main() {
		fruits := []string{"peach", "banana", "kiwi"}
		sort.Sort(ByLength(fruits))
		fmt.Println(fruits)
		// Output:[banana kiwi peach]
	}
			
