## func (b *Reader) ReadRune() (r rune, size int, err error)

返回值

- r 一个Unicode字符
- size r所占的字节数
- err 错误

功能说明

- ReadRune读取一个UTF-8编码的字符，并将其对应的Unicode编码和所占字节数返回。如果编码错误，ReadRune只读取一个字节并返回unicode.ReplacementChar(U+FFFD)和长度1。

代码示例

	package main

	import (
		"bufio"
		"bytes"
		"fmt"
	)

	func main() {
		b := []byte("你好世界")
		rb := bytes.NewBuffer(b)
		r := bufio.NewReader(rb)
		c, size, err := r.ReadRune()
		if err == nil {
			fmt.Println(string(c))
			fmt.Printf("%x, %d\n", c, size)
			fmt.Printf("%x\n", b[:size])
		}
	}

代码输出

	你
	4f60, 3
	e4bda0
