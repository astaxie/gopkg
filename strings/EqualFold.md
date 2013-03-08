# func EqualFold(s, t string) bool

参数列表

- s 表示需要判断的主串 
- t 表示需要比较的辅串

返回值：

- 返回bool

功能说明：

字符串s和t比较，它们在全部小写的情况下，采用UTF8编码的底层的unicode是否一致

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.EqualFold("Go", "go"))  //true
	}
