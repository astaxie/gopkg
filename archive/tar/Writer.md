## type Writer struct {}

功能说明：

该结构主要用于将文件写入到tar包中，并记录文件的数据头。

代码实例（本示例不包含目录递归功能，[完整实例](https://github.com/Unknwon/go-compresser/blob/master/go-tar.gz.go)）：

	package main
	
	import (
		"os"
		"io"
		"archive/tar"
	)
	
	func main() {
		fw, err := os.Create("demo.tar")	// 创建tar包文件，返回*io.Writer
		handleError(err)	// handleError为错误处理函数，下同
		defer fr.Close()
		
		// 实例化新的tar.Writer
		tw := tar.NewWriter(fw)
		defer tw.Close()
		
		// 获取要打包的文件的内容
		fr, err = os.Open(demo.txt)
		handleError(err)
		defer fr.Close()
		
		// 获取文件信息
		fi, err = fr.Sata()
		handleError(err)
		
		// 创建tar.Header
		hdr := new(tar.Header)
		hdr.Name = fi.Name()
		hdr.Size = fi.Size()
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()
		
		// 写入数据头
		err = tw.WriteHeader(hdr)
		handleError(err)
		
		// 写入文件数据
		_, err = io.Copy(tw, fr)
		handleError(err)
	}