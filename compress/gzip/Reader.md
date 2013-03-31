## type Reader struct {}

功能说明：

该结构主要用于读取gzip包的数据，但Go并没有提供像zip或tar那样便利的方法直接返回文件的切片，因此强烈建议和tar包一起使用！

代码实例（一般gzip与tar同时使用，具体方法请参考，[完整实例](https://github.com/Unknwon/go-compresser/blob/master/go-tar.gz.go)）：

	package main
	
	import (
		"os"
		"io"
		"archive/zip"
	)
	
	func main() {
		// 打开gzip文件
		fr, err := os.Open("demo.gzip")		
		handleError(err)	// handleError为错误处理函数，下同		defer fr.Close()
		
		// 创建gzip.Reader
		gr, err := gzip.NewReader(fr)
		handleError(err)
		defer gr.Close()
		
		// 读取文件内容
		buf := make([]byte, SIZE)	// 如果单独使用，需自己决定要读多少内容，根据官方文档的说法，你读出的内容可能超出你的所需（当你压缩多个文件时，再次强烈建议直接和tar组合使用）
		gr.Read(buf)
		
		// 将包中的文件数据写出
		fw, err := os.Create(gr.Header.Name)
		handleError(err)
		_, err = fw.Write(buf)
		handleError(err)
	}
