## func NewBuffer(buf []byte) *Buffer

参数列表

- buf 字节切片

返回值

- *Buffer

功能说明

- NewBuffer创建一个新的Buffer，并使用buf进行初始化。这个buf用来作为准备要读的数据；也可以用来指定写缓冲区的大小，这时buf应该是cap(buf)为指定大小，但是len(buf)为0。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		buf := []byte("123456")
		b := bytes.NewBuffer(buf)
		var data [6]byte
		b.Read(data[:])
		fmt.Println(string(data[:]))
	}

代码输出

	123456
