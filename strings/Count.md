# func Count(s, sep string) int

参数列表

- s 表示需要判断的主串 
- sep 需要计算的子串

返回值：

- 返回int 表示s中包含sep的个数 

功能说明：

这个函数主要是用来判断s中包含了多少个sep

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Count("cheese", "e"))  //3
		fmt.Println(strings.Count("cheese", "ee")) //1
		fmt.Println(strings.Count("five", ""))     // 5
	
	}