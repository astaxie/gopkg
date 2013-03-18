## func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

参数列表

- b 要存放读取数据的字节切片
- off 要读取数据的偏移位置

返回值

- n 已读取数据的字节数
- err 错误

功能说明

- ReadAt从Reader中偏移为off字节的位置读取len(b)个字节，并返回已读取的字节数；如果遇到错误则返回错误（通常是io.EOF）。ReadAt读取的数据不会从Reader中清除。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewReader([]byte("12345"))
		var buff [4]byte
		n, err := b.ReadAt(buff[:], 2)
		fmt.Println(n, err, string(buff[:n]))
		n, err = b.ReadAt(buff[:], 2)
		fmt.Println(n, err, string(buff[:n]))
		n, err = b.Read(buff[:])
		fmt.Println(n, err, string(buff[:n]))
	}

代码输出
	
	3 EOF 345
	3 EOF 345
	4 <nil> 1234
