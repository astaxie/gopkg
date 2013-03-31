# func ToTitle(s string) string

参数列表

- s 表示需要处理的字符串

返回值：

- 返回string 转化之后的字符串

功能说明：

该函数把s字符串里面的每个字符转化为首字母大写，其实和ToUpper一样的效果，但是有些语种的unicode,ToTitle和ToUpper效果不一样，但是我没试出来过，英语至少是一样的。

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.ToTitle("Gopher"))
		//GOPHER
	}