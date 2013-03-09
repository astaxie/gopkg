## type Writer struct {}

功能说明：

该结构主要用于将文件写入到zip包中，并记录每个文件的数据头。
另外要注意的是，Create()与CreateHeader()两个方法的效果是一样的，都是添加新的文件数据头到zip包文件中，并返回*io.Writer，区别在于前者指需要指定文件名，而后者需要传递一个完整的FileHeader结构作为参数。

代码实例（本示例不包含目录递归功能，[完整实例](https://github.com/Unknwon/go-compresser/blob/master/go-zip.go)）：

	package main
	
	import (
		"os"
		"io"
		"archive/tar"
	)
	
	func main() {
		fw, err := os.Create("demo.zip")	// 创建zip包文件，返回*io.Writer
		handleError(err)	// handleError为错误处理函数，下同
		defer fw.Close()
		
		// 实例化新的zip.Writer
		zw := zip.NewWriter(fw)
		defer zw.Close()
		
		// 获取要打包的文件的内容
		fr, err = os.Open("demo.txt")
		handleError(err)
		defer fr.Close()
		
		// 获取文件信息
		fi, err = fr.Sata()
		handleError(err)
		
		// 创建zip.FileHeader
		fh := new(zip.FileHeader)
		fh.Name = fi.Name()
		fh.UncompressedSize = uint32(fi.Size())
		fw, err := zw.CreateHeader(fh)
		handleError(err)

		// 读取文件数据
		buf := make([]byte, fi.Size())
		_, err = fr.Read(buf)
		handleError(err)

		// 写入数据到zip包
		_, err = fw.Write(buf)
		handleError(err)
	}