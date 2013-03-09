# func Index(s, sep string) int

参数列表

- s 表示需要判断的主串
- sep 需要判断的第一次出现位置的字符串

返回值：

- 返回int，对应sep出现在s中的位置

功能说明：

该函数主要判断sep串在s串中第一次出现的位置，如果不存在返回-1

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.Index("chickenkenken", "ken"))  //4
		fmt.Println(strings.Index("chicken", "dmr"))        //-1
	}