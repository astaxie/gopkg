# func (w *Writer) Reset(dst io.Writer)

参数列表：

- dst 重置时将为作w的下层io.Writer

功能说明：

Reset会丢弃当前的w的状态，这相当于把dst、w的级别和字典作为参数，重新调用NewWriter或者NewWriterDict函数一样。

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
	
		//reset
		buf1 := bytes.NewBuffer(nil)
		flateWrite.Reset(buf1)
		//写入待压缩内容
		flateWrite.Write([]byte("compress/flate\n"))
		flateWrite.Flush()
		fmt.Println(buf)  //什么都没有一个空行，因为我们reset重置了
		fmt.Println(buf1) //这个会输出结果，因为下层io.Writer被替换为buf1了
	}

