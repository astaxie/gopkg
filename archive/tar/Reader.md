## type Reader struct {}

功能说明：

该结构主要用于读取tar包时获取所有文件的slice，并调用Next方法逐个读取其中的文件。

代码实例（本示例不包含目录递归及目录判断功能，[完整实例](https://github.com/Unknwon/go-compresser/blob/master/go-tar.gz.go)）：

	package main
	
	import (
		"os"
		"io"
		"archive/tar"
	)
	
	func main() {
		fr, err := os.Open("dmeo.tar")	// 打开tar包文件，返回*io.Reader
		handleError(err)	// handleError为错误处理函数，下同
		defer fr.Close()
		
		// 实例化新的tar.Reader
		tr := tar.NewReader(fr)
		
		for	{
			hdr, err := tr.Next()	// 获取下一个文件，第一个文件也用此方法获取
			if err == io.EOF {
				break	// 已读到文件尾
			}
			handleError(err)
			
			// 通过创建文件获得*io.Writer
			fw, _ := os.Create("demo/" + hdr.Name)
			handleError(err)
			
			// 拷贝数据
			_, err = io.Copy(fw, tr)
			handleError(err)
		}
	}