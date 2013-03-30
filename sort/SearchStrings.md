## func SearchStrings(a []string, x string) int

参数列表

返回值：

- a 表示升序排列的 string 切片 
- x 表示搜索的 string 值 

功能说明：

SearchStrings 在 strings 切片中搜索x并返回索引,如 Search 函数所述. 返回可以插入 x 值的索引位置，如果 x 不存在，返回数组 a 的长度。切片必须以升序排列。
		
代码案例：
	
	package main
	
	import (
		"fmt"
		"sort"
	)
	
	func main() {
	
		a := []string{"Cupcake", "Donut", "Eclair", "Froyo","Gingerbread",
			"Honeycomb", "IceCreamSandwich", "JellyBean", "KeyLimePie"}
		
		fmt.Println(sort.SearchStrings(a, "Google")) // 5
	}
	