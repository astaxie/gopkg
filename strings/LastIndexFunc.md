# func LastIndexFunc(s string, f func(rune) bool) int

参数列表

- s 表示需要判断的主串
- f 表示一个函数，该函数参数是rune，返回值是bool，如果rune符合f函数的逻辑那么返回true，否者返回false

返回值：

- 返回int，对应符合函数f的字符的位置

功能说明：

该函数主要判断s中的每一个字符传入函数f，返回符合函数f的最后一个字符的位置，如果都不符合则返回-1

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.LastIndexFunc("astaxie", splitfunc)) //符合条件的有字符t和x，因为字符x是最后出现的位置，所以这个位置应该返回4
		fmt.Println(strings.LastIndexFunc("aaabbbb", splitfunc)) //所有的字符都不符合条件，则返回-1
	}
	
	func splitfunc(a rune) bool {
		if a > 'm' {
			return true
		}
		return false
	}
