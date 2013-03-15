## func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte

参数列表

- _case 转换规则
- s 要转换的字节切片

返回值

- []byte

功能说明

- ToUpperSpecial返回s的一个副本，并把其中的所有Unicode字符都根据_case指定的规则转换成大写。

代码示例

	package main

	import (
		"bytes"
		"fmt"
		"unicode"
	)

	func main() {
		s := []byte("Hello, world!")
		fmt.Println(string(bytes.ToUpperSpecial(unicode.AzeriCase, s)))
		fmt.Println(string(s))
	}

代码输出

	HELLO, WORLD!
	Hello, world!

