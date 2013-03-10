# func ToUpper(s string) string

参数列表

- s 表示需要处理的字符串

返回值：

- 返回string 转化之后的字符串

功能说明：

该函数把s字符串里面的每个字符转化为大写

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.ToUpper("Gopher"))
		//GOPHER
	}