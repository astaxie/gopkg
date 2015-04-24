# func (w *Writer) Flush() error

返回值：返回一个error，没有错误时该error为nil

功能说明：

Flush将缓存中的压缩数据刷新到下层的io.writer中。它主要用在压缩的网络协议中，目的是确保远程读取器有足够的数据重建一个数据包。Flush是阻塞的，直到缓冲中的数据都被写入到下层io.writer中才返回。如果下层io.writer返回一个error，那么Flush也会返回该error。

在zlib库的术语中，Flush等同于Z_SYNC_FLUSH。

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
