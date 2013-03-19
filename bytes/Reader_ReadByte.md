## func (r *Reader) ReadByte() (b byte, err error)

返回值

- b 读取的字节
- err 错误

功能说明

- ReadByte从Reader中读取一个字节并返回，如果遇到错误则返回错误（通常是io.EOF）。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewReader([]byte("12345"))
		c, err := b.ReadByte()
		fmt.Println(string(c), err)
	}

代码输出
	
	1 <nil>
