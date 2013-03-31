## func (r *Reader) ReadRune() (ch rune, size int, err error)

返回值

- ch 读取的Unicode字符
- size ch的UTF-8编码的字节数
- err 错误

功能说明

- ReadRune从Reader中按照UTF-8编码读取一个Unicode字符，返回此字符，其字符的UTF-8编码占用的字节数，如果遇到错误则返回错误（通常是io.EOF）。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewReader([]byte("你好世界"))
		r, size, err := b.ReadRune()
		fmt.Println(string(r), size, err)
	}

代码输出
	
	你 3 <nil>
