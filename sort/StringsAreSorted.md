## func StringsAreSorted(a []string) bool

参数列表

- a 表示要判断的 string 切片

功能说明：

StringsAreSorted 判断 string 切片是否已经按升序排列

代码案例：
	
	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
		a := []string{"PHP", "golang", "java", "python", "C", "Objective-C"}
		fmt.Println(sort.StringsAreSorted(a)) // false
		sort.Strings(a)
		fmt.Println(sort.StringsAreSorted(a)) //true
	}
	