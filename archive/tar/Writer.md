## type Writer struct {}

功能说明：

该结构主要用于将文件写入到tar包中，并记录每个文件的数据头。
另外要注意的是，本例所采用的是常规的方法，由于Go1.0.3版本仅允许一般的头文件名（Header.Name）总长度为100个字节，因此无法做很多层目录的递归（可以做，但文件名的总长度受限），如果您知道如何将该长度扩展至更多字节，请不吝指出，以帮助大家。

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
		defer fw.Close()
		
		// 实例化新的tar.Writer
		tw := tar.NewWriter(fw)
		defer tw.Close()
		
		// 获取要打包的文件的内容
		fr, err = os.Open("demo.txt")
		handleError(err)
		defer fr.Close()
		
		// 获取文件信息
		fi, err = fr.Stat()
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
