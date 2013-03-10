# func ContainsAny(s, chars string) bool

参数列表

- s 表示需要判断的主串 
- chars 表示保存的unicode字符串

返回值：

- 返回bool

功能说明：

这个函数主要是用来判断s中是否包含chars中的字符中的任意字符，如果包含返回true，否者返回false

代码实例：

  package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.ContainsAny("team", "i"))       //false
		fmt.Println(strings.ContainsAny("failure", "wwwi")) //true
		fmt.Println(strings.ContainsAny("foo", ""))         //false
		fmt.Println(strings.ContainsAny("", ""))            //false
	}
