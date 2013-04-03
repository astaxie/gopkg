# func HasSuffix(s, suffix string) bool

参数列表

- s 表示需要判断的主串
- suffix 需要判断的后缀字符串

返回值：

- 返回bool，

功能说明：

该函数主要判断s串中是否含有后缀suffix，如果包含，那么返回true，否则返回false

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.HasSuffix("astaxie", "as")) //false
		fmt.Println(strings.HasSuffix("astaxie", "xie")) //true
	}