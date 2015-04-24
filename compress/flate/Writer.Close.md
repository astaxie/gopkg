# func (w *Writer) Close() error

返回值：返回一个error，没有错误时该error为nil

功能说明：

刷新缓冲并关闭w

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
		//刷新缓存并关闭flateWrite
		defer flateWrite.Close()
		//写入待压缩内容
		flateWrite.Write([]byte("compress/flate\n"))
		flateWrite.Flush()
		fmt.Println(buf)
	
	}
