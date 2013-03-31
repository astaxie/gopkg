# func LastIndex(s, sep string) int

参数列表

- s 表示需要判断的对象字符串
- sep 表示最后出现的字符串

返回值：

- 返回int

功能说明：

该函数主要判断sep串在s串中最后一次出现的位置，如果不存在返回-1 

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		fmt.Println(strings.LastIndex("chickenkenken", "ken"))  //10
		fmt.Println(strings.LastIndex("chicken", "dmr"))        //-1
	}