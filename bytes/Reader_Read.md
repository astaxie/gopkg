## func (r *Reader) Read(b []byte) (n int, err error)

参数列表

- b 要存放读取数据的字节切片

返回值

- n 已读取的字节数
- err 错误

功能说明

- Read从Reader中读取len(b)个字节，返回已读取的字节数，如果遇到错误则返回错误（通常是io.EOF）。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewReader([]byte("12345"))
		var buff [3]byte
		n, err := b.Read(buff[:])
		fmt.Println(n, err, string(buff[:n]))
		n, err = b.Read(buff[:])
		fmt.Println(n, err, string(buff[:n]))
		n, err = b.Read(buff[:])
		fmt.Println(n, err, string(buff[:n]))
	}

代码输出
	
	3 <nil> 123
	2 <nil> 45
	0 EOF
