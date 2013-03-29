## func Search(n int, f func(int) bool) int 

参数列表

- n 表示切片的长度
- f func(int) bool类型的函数

返回值：

- 返回切片的索引

功能说明：
		
Search 使用二分查找法在 [0, n) 中寻找并返回满足 f(i)==true 的最小索引 i ，假定该索引在区间 [0, n)内，则	f(i) == true 就蕴含f(i+1) == true。Search 要求 f 对于输入区间 [0, n)（可能为空）的前一部分为false，而对于剩余（可能为空）的部分为 true；

Search 返回第一个 f 为 true 时的索引 i。若该索引不存在，Search 就返回 n。Search仅当 i 在区间 [0, n) 内时才调用 f(i)。
		
Search 常用于在一个已排序的，可索引的数据结构中寻找索引为 i 的值 x，例如数组或切片。这种情况下，实参f，一般是一个闭包，会捕获所要搜索的值，以及索引并排序该数据结构的方式。
		
例如，给定一个以升序排列的切片数据，调用
		
		Search(len(data), func(i int) bool {
				return data[i] >= 23
		})
		
会返回满足 data[i] >= 23 的最小索引 i。若调用者想要判断 23 是否在此切片中，就必须单独测试 data[i] == 23的值。
		
搜索降以序排列的数据，需使用 <= 操作符，而非 >= 操作符。	
		
代码案例：

====
	package main
	
	import (
		"fmt"
		"sort"
	)
	
	type finder struct {
		data []int
		targ int
		f    func(n int) bool
	}
	
	func MakeFinder1() *finder {
		var f finder
		f.f = func(i int) bool {
			return f.data[i] >= f.targ
		}
		return &f
	}
	
	func MakeFinder2() *finder {
		var f finder
		f.f = func(i int) bool {
			return f.data[i] <= f.targ
		}
		return &f
	}
	
	func (f *finder) Find(data []int, x int) int {
		f.data = data
		f.targ = x
		return sort.Search(len(f.data), f.f)
	}
	
	func main() {
		var data1 = []int{9, 19, 29, 39, 49, 59, 69, 79, 89, 99}
		var data2 = []int{99, 89, 79, 69, 59, 49, 40, 39, 29, 19, 9}
		f1 := MakeFinder1()
		i := f1.Find(data1, 50)
		fmt.Println(i, data1[i]) // 5 59
	
		f2 := MakeFinder2()
		i = f2.Find(data2, 40)
		fmt.Println(i, data2[i]) // 6 40
		if i < len(data2) && f2.targ == data2[i] {
			fmt.Printf("find %v in data2", f2.targ)
		} else {
			fmt.Printf("can'tfind %v in data2", f2.targ)
		}
		// find 40 in data2 
	}

