# func NewWriter(w io.Writer, level int) (*Writer, error)

参数列表；

* w 输出数据的Writer
* level 压缩级别

返回列表：

* *Writer 基于压缩级别新生成的压缩数据的Writer
* error 该函数的错误信息

功能说明：

该函数返回一个压缩级别为level的新的压缩用的Writer。压缩级别的范围是1 (BestSpeed) to 9 (BestCompression)。压缩效果越好的意味着压缩速度越慢。0 (NoCompression)表示不做任何压缩；仅仅只需要添加必要的DEFLATE信息。-1 (DefaultCompression)表示用默认的压缩级别。
如果压缩级别在-1到9的范围内，error返回nil，否则将返回非nil的错误信息

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
		flateWrite, err := flate.NewWriter(buf, flate.BestCompression)
		if err != nil {
			log.Fatalln(err)
		}
		defer flateWrite.Close()
		//写入待压缩内容
		flateWrite.Write([]byte("compress/flate\n"))
		flateWrite.Flush()
		fmt.Println(buf)
	
	}