# func ToLowerSpecial(_case unicode.SpecialCase, s string) string

参数列表

- _case 表示unicode的SpecialCase对象
- s 表示需要处理的字符串

返回值：

- 返回string 转化之后的字符串

功能说明：

该函数把s字符串里面的每个单词转化为小写，但是调用的是unicode.SpecialCase的ToLower方法

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
		"unicode"
	)
	
	func main() {
		var SC unicode.SpecialCase
		fmt.Println(strings.ToLowerSpecial(SC, "Gopher"))
		//gopher
	}
