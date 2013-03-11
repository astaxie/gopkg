## func NewReadWriter(r *Reader, w *Writer) *ReadWriter

参数列表
- r 表示要读取的来源
- w 表示要写入的目的

返回值:

- 返回一个可读写的对象

功能说明：

这个函数根据指定的可读对象和可写对象创建并返回一个可读可写的对

代码示例：

	package main

	import (
		"bufio"
		"bytes"
		"fmt"
	)

	func main() {
		rb := bytes.NewBuffer([]byte("a string to be read"))
		wb := bytes.NewBuffer(nil)
		r := bufio.NewReader(rb)
		w := bytes.NewBuffer(wb)
		rw := bufio.NewReadWriter(r, w)
		// use rw to read
		var rbuf [128]byte
		if n, err := rw.Read(rbuf[:]); err != nil {
			return
		} else {
			fmt.Println(string(rbuf[:n]))
		}
		// use rw to write
		rw.Write([]byte("a string to be written"))
		rw.Flush()
		fmt.Println(string(wb.Bytes()))
	}

代码输出：

	a string to be read
	a string to be written
