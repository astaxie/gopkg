## func Runes(s []byte) []rune

参数列表

- s 字节切片

返回值

- []rune s对应的Unicode字符切片

功能说明

- Runes把s解释为UTF-8编码的字节序列，并返回对应的Unicode切片。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("你好世界")
		fmt.Printf("%d, %s\n", len(s), string(s))
		r := bytes.Runes(s)
		fmt.Printf("%d, %s\n", len(r), string(r))
	}

代码输出

	12, 你好世界
	4, 你好世界
