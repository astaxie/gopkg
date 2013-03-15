## func Join(a [][]byte, sep []byte) []byte

参数列表

- a 包含要被连接的字节切片的切片
- sep 用于连接的字节切片

返回值

- []byte 连接后的字节切片

功能说明

- Join用sep把a中的每个字节切片连接成一个字节切片并返回。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		a := [][]byte{
			[]byte("hello"),
			[]byte("world"),
		}
		fmt.Println(string(bytes.Join(a, []byte(", "))))
		fmt.Println(string(bytes.Join(a, []byte{})))
		fmt.Println(string(bytes.Join(a, nil)))
	}

代码输出

	hello, world
	helloworld
	helloworld
