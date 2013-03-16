## func (b *Buffer) WriteRune(r rune) (n int, err error)

参数列表

- r 要写入的Unicode字符

返回值

- n 写入的字节数
- err 错误

功能说明

- WriteRune把一个Unicode字符的UTF-8编码写入Buffer，并返回写入字节数。err总是nil，这个返回类型是为了和bufio.Writer的WriteRune函数匹配。如果Buffer的缓冲区太大，则WriteRune会panic并产生ErrTooLarge异常。

代码示例

  	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer(nil)
		b.WriteRune('你')
		b.WriteRune('好')
		fmt.Println(string(b.Bytes()))
	}
	
代码输出

	你好
