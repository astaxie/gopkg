## func NewReader(rd io.Reader) *Reader

参数列表

- rd io.Reader接口，创建的Reader对象会从此接口读取数据

返回值：

- *Reader

功能说明：

- 创建支持缓存读取的具有缺省长度缓冲区的Reader对象，Reader对象会从底层的io.Reader接口读取尽量多的数据进行缓存。

代码示例：

	package main

	import (
		"bytes"
		"bufio"
		"fmt"
	)

	func main() {
		rb := bytes.NewBuffer([]byte("a string to be read"))
		r := bufio.NewReader(rb)
		var buf [128]byte
		n, _ := r.Read(buf[:])
		fmt.Println(string(buf[:n]))
	}

代码输出：

	a string to be read
