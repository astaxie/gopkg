# func TrimSpace(s string) string

参数列表

- s 表示需要处理的字符串

返回值：

- 返回string 转化之后的字符串

功能说明：

该函数把s字符串开头或者结尾里面空白符('\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP))全部过滤掉，返回过滤之后的字符串

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
		//a lone gopher
	}