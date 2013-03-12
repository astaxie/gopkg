## func NewWriter(wr io.Writer) *Writer

参数列表

- wr io.Writer接口，创建的Writer对象会将数据写入此接口

返回值：

- *Writer

功能说明：

- 创建支持缓存写的具有缺省长度缓冲区的Writer对象，Writer对象会将缓存的数据批量写入底层的io.Writer接口

代码示例：

	package main

	import (
		"bytes"
		"bufio"
		"fmt"
	)

	func main() {
		wb := bytes.NewBuffer(nil)
		w := bufio.NewWriter(wb)
		w.Write([]byte("hello,"))
		w.Write([]byte("world!"))
		fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
		w.Flush()
		fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
	}

代码输出：

	0:
	 12:hello,world! 
