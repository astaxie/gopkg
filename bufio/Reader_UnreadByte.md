## func (b *Reader) UnreadByte() error

返回值

- error 错误

功能说明

- UnreadByte取消已读取的最后一个字节（即把字节重新放回读取缓冲区的前部）。只有最近一次读取的单个字节才能取消读取。

代码示例

	package main

	import (
		"bufio"
		"bytes"
		"fmt"
	)

	func main() {
		rb := bytes.NewBuffer([]byte("1234,56$"))
		r := bufio.NewReader(rb)
		line, _ := r.ReadSlice(',')
		fmt.Println(string(line))
		// unread ','
		fmt.Println(r.UnreadByte())
		line, _ = r.ReadSlice('$')
		fmt.Println(string(line))
	}

代码输出

	1234,
	<nil>
	,56$
