## func (b *Buffer) ReadRune() (r rune, size int, err error)

返回值

- r Unicode字符
- size r的UTF-8编码占用的字节数
- err 错误

功能说明

- ReadRune从Buffer中读取一个Unicode字符，同时返回它的UTF-8编码的字节数；如果没有数据可读则返回io.EOF。如果读取的不是UTF-8编码的字节序列，则读取一个字节并返回U+FFFD。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("你好世界"))
		r, size, err := b.ReadRune()
		fmt.Println(string(r), size, err)
	}
	
代码输出

	你 3 <nil>
