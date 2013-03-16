## func (b *Buffer) UnreadByte() error

返回值

- err 错误

功能说明

- UnreadByte取消上次读取的最后一个字节。如果上次读取之后有过写操作，则返回错误。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("12345"))
		fmt.Println(string(b.Next(3)))
		b.UnreadByte()
		fmt.Println(b.String())
	}
	
代码输出

	123
	345
