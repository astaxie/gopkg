# func TrimFunc(s string, f func(rune) bool) string

参数列表

- s 表示需要处理的字符串
- f 表示一个函数，判断是否过滤该字符的函数，如果为真那么就过滤

返回值：

- 返回string 转化之后的字符串

功能说明：

该函数把s字符串里面开头和结尾部分字符传入f函数进行判断是否过滤，为真就过滤

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Printf("[%q]", strings.TrimFunc("xxastaxieyy", filte))
		//astaxie
	}
	
	func filte(r rune) bool {
		if r > 't' {
			return true
		}
		return false
	}
