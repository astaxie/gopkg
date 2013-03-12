## func NewReadWriter(r *Reader, w *Writer) *ReadWriter

参数列表
- r 要读取的来源Reader
- w 要写入的目的Writer

返回值:

- *ReadWriter

功能说明：

- 这个函数根据指定的Reader和Writer创建一个ReadWriter对象

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
		if _, err := rw.Write([]byte("a string to be written")); err != nil {
			return
		}
		rw.Flush()
		fmt.Println(string(wb.Bytes()))
	}

代码输出：

	a string to be read
	a string to be written
