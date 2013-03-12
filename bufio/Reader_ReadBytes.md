## func (b *Reader) ReadBytes(delim byte) (line []byte, err error)

参数列表

- delim 分隔字节

返回值

- line 读取的字节切片
- err 错误

功能说明

- ReadBytes读取数据直到delim第一次出现，返回读取的字节序列（包括delim）。如果ReadBytes在读到第一个delim之前出错，它返回已读取的数据和那个错误（通常是io.EOF）。只有当返回的数据不以delim结尾时，返回的err才不为空值。

代码示例

	package main

	import (
		"bytes"
		"bufio"
		"fmt"
	)

	func main() {
		rb := bytes.NewBuffer([]byte("123$456"))
		r := bufio.NewReader(rb)
		b, err := r.ReadBytes('$')
		fmt.Printf("%s, %v\n", string(b), err)
	}

代码输出

	123$, <nil>
