# func Replace(s, old, new string, n int) string

参数列表

- s 表示需要替换的主字符串
- old 表示需要替换的字符串
- new 表示替换的新字符串
- n 表示替换的次数

返回值：

- 返回string，处理之后的字符串

功能说明：

该函数实现在s中把old替换为new字符串，替换次数为n，如果n小于0，那么就全部替换

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))   //oinky oinky oink
		fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))   //moo moo moo
	}