## func (b *Buffer) ReadByte() (c byte, err error)

返回值

- c 读取的字节
- err 错误

功能说明

- ReadByte读取一个字节返回，如果没有数据可读则err为io.EOF。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("123456"))
		c, err := b.ReadByte()
		fmt.Println(c, err)
		fmt.Println(string(b.Bytes()))
	}
	
代码输出

	49 <nil>
	23456
