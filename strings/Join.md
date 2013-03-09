# func Join(a []string, sep string) string

参数列表

- a 表示需要链接起来的字符串slice
- sep 表示链接的符号

返回值：

- 返回string

功能说明：

该函数主要实现字符串slice的链接功能，把slice的每一个元素通过sep进行链接

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		s := []string{"foo", "bar", "baz"}
		fmt.Println(strings.Join(s, ", ")) //foo, bar, baz
	}