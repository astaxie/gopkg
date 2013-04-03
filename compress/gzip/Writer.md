## type Writer struct {}

功能说明：

该结构主要用于将文件写入到gzip包中，并记录每个文件的数据头。
另外要注意的是，NewWriter()与NewWriterLevel()的主要区别在于前者使用默认的压缩级别，而后者可以指定自定义的压缩级别。单独操作gzip很蛋疼，建议直接与tar结合将会十分方便。

代码实例（一般gzip与tar同时使用，具体方法请参考，[完整实例](https://github.com/Unknwon/go-compresser/blob/master/go-tar.gz.go)）：

	package main
	
	import (
		"os"
		"io"
		"compress/gzip"
	)
	
	func main() {
		fw, err := os.Create("demo.gzip")	// 创建gzip包文件，返回*io.Writer
		handleError(err)	// handleError为错误处理函数，下同
		defer fw.Close()
		
		// 实例化新的gzip.Writer
		gw := gzip.NewWriter(fw)
		defer gw.Close()
		
		// 获取要打包的文件的内容
		fr, err = os.Open("demo.txt")
		handleError(err)
		defer fr.Close()
		
		// 获取文件信息
		fi, err = fr.Sata()
		handleError(err)
		
		// 创建gzip.Header
		gw.Header.Name = fi.Name()

		// 读取文件数据
		buf := make([]byte, fi.Size())
		_, err = fr.Read(buf)
		handleError(err)

		// 写入数据到zip包
		_, err = fw.Write(buf)
		handleError(err)
	}
