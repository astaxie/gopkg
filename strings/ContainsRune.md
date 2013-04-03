# func ContainsRune(s string, r rune) bool

参数列表

- s 表示需要判断的主串 
- r 表示rune字符

返回值：

- 返回bool

功能说明：

这个函数主要是用来判断s中是否包含rune类型的r字符，如果包含返回true，否者返回false

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.ContainsRune("team", rune('m')))    //true
		fmt.Println(strings.ContainsRune("failure", rune('w'))) //false
		fmt.Println(strings.ContainsRune("谢foo", rune('谢')))    //true
		fmt.Println(strings.ContainsRune("", 30))               //false
	}

