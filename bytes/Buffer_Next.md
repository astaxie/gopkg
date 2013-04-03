## func (b *Buffer) Next(n int) []byte

参数列表

- n 要读取的字节数

返回值

- []byte 字节切片

功能说明

- Next返回Buffer中下n个未读取的字节并返回，本次读取会修改Buffer中的读取位置。如果Buffer中未读数据不足n自己，则返回全部未读数据。返回的字节切片在下次读或者写前有效。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("123456"))
		fmt.Println(string(b.Next(4)))
		fmt.Println(string(b.Next(4)))
	}
	
代码输出

	1234
	56
