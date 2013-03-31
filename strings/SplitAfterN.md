# func SplitAfterN(s, sep string, n int) []string

参数列表

- s 表示需要处理的字符串
- sep 表示分割的字符串
- n 表示需要分割的子字符串

	- n > 0: 最多n个子字符串; 最后一个就是剩下未分割的子字符串.
	- n == 0: 返回为0的字符串
	- n < 0: 返回所有的子字符串，和SplitAfter

返回值：

- 返回[]string 分割之后的字符串slice

功能说明：

该函数s根据sep分割，返回分割之后子字符串的slice,和split一样，只是返回的子字符串保留sep，如果sep为空，那么每一个字符都分割

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))  //["a," "b,c"]
		fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", "", 5))   //["a" "," "b" "," "c"]
		fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 1))   //["a,b,c"]
	}