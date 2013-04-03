# func Split(s, sep string) []string

参数列表

- s 表示需要处理的字符串
- sep 表示分割的字符串

返回值：

- 返回[]string 分割之后的字符串slice

功能说明：

该函数s根据sep分割，返回分割之后子字符串的slice，如果sep为空，那么每一个字符都分割

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Printf("%q\n", strings.Split("a,b,c", ","))   //["a" "b" "c"]
		fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))  //["" "man " "plan " "canal panama"]
		fmt.Printf("%q\n", strings.Split(" xyz ", ""))  //[" " "x" "y" "z" " "]
		fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))   //[""]
	}
