## func NewBufferString(s string) *Buffer

参数列表

- s 用于初始化Buffer内容的字符串

返回值

- *Buffer

功能说明

- NewBufferString创建一个Buffer，并用s初始化它的内容。它用于读已经存在的数据。大多数情况下，new(Buffer)或者直接声明一个Buffer变量就可以了。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		b := bytes.NewBufferString("123456")
		var data [6]byte
		b.Read(data[:])
		fmt.Println(string(data[:]))
	}

代码输出

	123456
