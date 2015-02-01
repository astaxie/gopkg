# func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser

参数列表

- r DEFLATE压缩的数据
- dict 解压数据时预设的字典，和NewWriterDict函数里的dict相同

返回值：解压后的ReadCloser数据 

功能说明：

从r读取DEFLATE压缩数据，使用预设的dict字典解压数据，返回一个解压过的io.ReadCloser，使用后需要调用者关闭该io.ReadCloser。主要用来读取NewWriterDict压缩的数据

示例：

	package main
	
	import (
		"bytes"
		"compress/flate"
		"fmt"
		"io"
		"log"
		"os"
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
		flateWrite.Write([]byte("compress/flate\n"))
		flateWrite.Flush()
		fmt.Println(buf)
	
		//解压刚刚压缩的内容
		flateReader := flate.NewReaderDict(buf, []byte("key"))
		//flateReader := flate.NewReader(buf)
		defer flateReader.Close()
		//输出
		io.Copy(os.Stdout, flateReader)
	}
