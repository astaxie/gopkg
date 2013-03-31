## func (b *Buffer) WriteByte(c byte) error

参数列表

- c 要写入的字节

返回值

- err 错误

功能说明

- WriteByte写入一个字节，总是返回nil。这个返回值只是为了匹配bufio.Writer的Write函数。如果内部数据太大，WriteByte会panic并产生ErrTooLarge异常。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer(nil)
		b.WriteByte('1')
		b.WriteByte('2')
		b.WriteByte('3')
		fmt.Println(string(b.Bytes()))
	}
	
代码输出

	123
