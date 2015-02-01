# func (w *Writer) Write(data []byte) (n int, err error)

参数列表：

- data 要写入的字节数据

返回值：

- n 写入的字节数
- err 错误信息，无错误返回nil

功能说明：

Write向w写入数据，最终会将压缩格式的数据写入到w的下层io.Writer中

示例：

	package main
	
	import (
		"bytes"
		"compress/flate"
		"fmt"
		"log"
	)
	
	func main() {
		//一个缓冲区存储压缩的内容
		buf := bytes.NewBuffer(nil)
	
		//创建一个flate.Writer
		flateWrite, err := flate.NewWriterDict(buf, flate.BestCompression, []byte("key"))
		if err != nil {
			log.Fatalln(err)
		}
		defer flateWrite.Close()
	
		//写入待压缩内容
		n, _ := flateWrite.Write([]byte("compress/flate\n"))
		flateWrite.Flush()
		fmt.Println(n) //15
	}
