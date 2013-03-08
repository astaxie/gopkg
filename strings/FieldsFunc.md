# func FieldsFunc(s string, f func(rune) bool) []string

参数列表

- s 表示需要判断的主串
- f 表示一个函数，该函数 

返回值：

- 返回[]string 

功能说明：

s按照一个空格或者多个连续的空格分割，返回分割之后的串数组，如果字符串没有空格或者只有一个空格，那么返回元素为去空格的s字符串

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) //Fields are: ["foo" "bar" "baz"]
		fmt.Printf("Fields are: %q", strings.Fields(" baz ")) //Fields are: ["bar"]
	}
