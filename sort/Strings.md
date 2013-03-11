## func Strings(a []string)

参数列表

- a 表示要排序的 string 切片

功能说明：

Strings 以升序排列 string 切片

代码实例：

	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		s := []string{"PHP", "golang", "java", "python", "C", "Objective-C"}
		sort.Strings(s)
		fmt.Println(s) // [C Objective-C PHP golang java python]
	}
	
