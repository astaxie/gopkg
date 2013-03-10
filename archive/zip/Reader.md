## type Reader struct {}

功能说明：

该结构主要用于读取zip包时获取所有文件的slice，并使用for循环逐个读取其中的文件。
另外要注意的是，这里的ReadCloser用于读取zip的包文件，而Reader则是读取包中每个单独的文件。

代码实例（本示例不包含目录递归及目录判断功能，[完整实例](https://github.com/Unknwon/go-compresser/blob/master/go-zip.go)）：

	package main
	
	import (
		"os"
		"io"
		"archive/zip"
	)
	
	func main() {
		// 用zip.OpenReader打开zip包文件
		rc, err := zip.OpenReader("demo.zip")
		handleError(err)	// handleError为错误处理函数，下同		defer rc.Close()
		
		// 使用循环逐个读取zip包内的单独文件
		for _, f := range rc.File {
			// 打开包中的文件
			r, err := f.Open()
			handleError(err)
			
			// 将包中的文件数据写出
			fw, _ := os.Create(f.FileInfo().Name())
			handleError(err)
			_, err = io.Copy(fw, r)
			handleError(err)
		}
	}