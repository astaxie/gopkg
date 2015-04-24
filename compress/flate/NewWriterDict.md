# func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)

参数列表；

* w 输出数据的Writer
* level 压缩级别
* dict 压缩预设字典

返回列表：

* *Writer 基于压缩级别和预设字典新生成的压缩数据的Writer
* error 该函数的错误信息

功能说明：

该函数和NewWriter差不多，只不过使用了预设字典进行初始化Writer，使用该Writer压缩的数据只能被使用同样字典初始化的Reader解压。可以实现基于密码的解压缩。

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
	
		//创建一个flate.Writer，压缩级别最好
		flateWrite, err := flate.NewWriterDict(buf, flate.BestCompression,[]byte("key"))
		if err != nil {
			log.Fatalln(err)
		}
		defer flateWrite.Close()
		//写入待压缩内容
		flateWrite.Write([]byte("compress/flate\n"))
		flateWrite.Flush()
		fmt.Println(buf)
	
	}