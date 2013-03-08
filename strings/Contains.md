## func Contains(s, substr string) bool

参数列表

- s 表示需要判断的主串 
- substr 表示包含的子串

返回值：

- 返回bool

功能说明：

这个函数主要是用来判断s中是否包含substr这个子串，如果包含返回true，否者返回false

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Contains("seafood", "foo")) //true
		fmt.Println(strings.Contains("seafood", "bar")) //false
		fmt.Println(strings.Contains("seafood", ""))    //true
		fmt.Println(strings.Contains("", ""))           //true
	}

