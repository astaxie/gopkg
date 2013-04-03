## func (b *Reader) UnreadRune() error

返回值

- error 错误

功能说明

- UnreadRune取消读取最后一次读取的Unicod字符。如果最后一次读取操作不是ReadRune，UnreadRune会返回一个错误（在这方面它比UnreadByte更严格，因为UnreadByte会取消上次任意读操作的最后一个字节）。

代码示例

	package main

	import (
		"bufio"
		"bytes"
		"fmt"
	)
	
	func main() {
		rb := bytes.NewBuffer([]byte("123456"))
		r := bufio.NewReader(rb)
		r.ReadByte()
		// error occurs
		fmt.Println(r.UnreadRune())
		c, _, _ := r.ReadRune()
		fmt.Printf("read %s\n", string(c))
		// no error happens
		fmt.Println(r.UnreadRune())
		c, _, _ = r.ReadRune()
		fmt.Printf("read %s\n", string(c))
	}

代码输出

	bufio: invalid use of UnreadRune
	read 2
	<nil>
	read 2
