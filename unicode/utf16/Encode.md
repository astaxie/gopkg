# func Encode(s []rune) []utf16 

参数列表: 

-s Unicode 编码序列 

返回值列表:

- 返回utf16序列串

功能说明：

>将 s 编码成 UTF-16 序列并返回


实例代码:
	package main

	import (
	 "fmt"
 	"unicode/utf16"
	)
	func main() {
	s := []rune("Hello 世界")
	u := utf16.Encode(s)
	fmt.Printf("%v", u)
	// [72 101 108 108 111 32 19990 30028]
	}

