## func (b *Buffer) Read(p []byte) (n int, err error)

参数列表

- p 要存放读取数据的字节切片

返回值

- n 已读取的字节数
- err 错误

功能说明

- Read从Buffer中读取len(p)个字节复制到p中；如果Buffer中未读数据不足len(p)则把所有的数据都复制到p中。返回实际读取的字节数。如果Buffer中没有数据，则err为io.EOF（除非len(p)为0）；否则err为nil。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("123456"))
		n, err := b.Read(nil)
		fmt.Printf("%d, %v\n", n, err)
		var buff [6]byte
		n, err = b.Read(buff[:])
		fmt.Printf("%d, %v, %s\n", n, err, buff[:n])
		n, err = b.Read(buff[:])
		fmt.Printf("%d, %v, %s\n", n, err, buff[:n])
	}
	
代码输出

	0, <nil>
	6, <nil>, 123456
	0, EOF,
