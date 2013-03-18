## func (r *Reader) UnreadByte() error

返回值

- error 错误

功能说明

- UnreadByte取消上次读取的最后一个字节，如果Buffer没有被读取过则UnreadByte也返回一个错误。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewReader([]byte("12345"))
		fmt.Println(b.UnreadByte())
		var buff [6]byte
		n, _ := b.Read(buff[:])
		fmt.Printf("read: %s\n", string(buff[:n]))
		fmt.Println(b.UnreadByte())
		n, _ = b.Read(buff[:])
		fmt.Printf("read: %s\n", string(buff[:n]))
	}

代码输出
	
	bytes.Reader: at beginning of slice
	read: 12345
	<nil>
	read: 5
