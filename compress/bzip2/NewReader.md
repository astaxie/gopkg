# func NewReader(r io.Reader) io.Reader

参数列表：

- r bzip2压缩数据

返回值：解压后的数据

功能说明：

返回一个从r读取bzip2压缩数据，然后返回一个解压后io.Reader

示例：

	package main
	
	import (
		"archive/tar"
		"compress/bzip2"
		"fmt"
		"io"
		"os"
		"path"
	)
	
	func main() {
		//打开一个bz2压缩文件
		bzip2File, _ := os.Open("demo.tar.bz2")
		defer bzip2File.Close()
	
		//进行解压
		r := bzip2.NewReader(bzip2File)
		//使用tar读取输出文件
		tr := tar.NewReader(r)
		for {
			tarHead, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("the tar file err is ", err)
				continue
			}
			fmt.Println(tarHead.Name)
	
			os.MkdirAll(path.Dir(tarHead.Name), os.ModePerm)
			fw, _ := os.Create(tarHead.Name)
			defer fw.Close()
	
			_, err = io.Copy(fw, tr)
			if err != nil {
				fmt.Println("copy file err is ", err)
				continue
			}
		}
	}
