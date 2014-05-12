# func DecodeRune(r1, r2 runc) rune

参数列表:

- r1 是utf-16代理对的高位码元
- r2 是utf-16代理对的低位码元


返回值列表:

- 返回值为解码后的 Unicode字符
- 如果r1 或者 r2不是有效的utf-16代理区字符
- 则返回 U+FFFD

功能说明:

>将utf-16代理对解码成一个Unicode字符

实例代码:

	package main

	import (
	"fmt"
	"unicode/utf16"
	)
	func main() {
	r := utf16.DecodeRune(0xDBFC, 0xDC00)
	fmt.Printf("%U", r)
	// U+10F000
	}




