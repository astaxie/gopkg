# func Decode(s []uint16) []rune

函数列表

- s utf-16序列 

返回值:

- 成功返回 Unicode 字符串

功能说明:

>将utf-16序列 解码成Unicode字符序列并返回

代码实例:

	package main

	import (
	 "fmt"
	"unicode/utf16"
	)
	func main() {
	u := []uint16{72, 101, 108, 108, 111, 32, 19990, 30028}
	s := utf16.Decode(u)
	fmt.Printf("%c", s)
	// [H e l l o   世 界]
	}



