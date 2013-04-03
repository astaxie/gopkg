# func IndexRune(s string, r rune) int

参数列表

- s 表示需要判断的主串
- r 需要判断的第一次出现位置的unicode码
返回值：

- 返回int，对应r出现在s中的位置

功能说明：

该函数主要判断unicode r在s串中第一次出现的位置，如果不存在返回-1

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.IndexRune("chicken", 'k'))  //4
		fmt.Println(strings.IndexRune("chicken", 'd'))  //-1
	}
