## func Fields(s []byte) [][]byte

参数列表

- s 要拆分的字节切片

返回值

- [][]byte 已拆分的字节切片的切片

功能说明

- Fields把字节切片s按照一个或者连续多个空白字符分隔成多个字节切片，如果s只包含空白字符则返回空字节切片

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("a b\tc\rd\ne   f")
		for i, f := range bytes.Fields(s) {
			fmt.Printf("[%d]%s, %d\n", i, string(f), len(f))
		}
	}

代码输出

	[0]a, 1
	[1]b, 1
	[2]c, 1
	[3]d, 1
	[4]e, 1
	[5]f, 1
