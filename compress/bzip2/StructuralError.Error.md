# func (s StructuralError) Error() string

返回值：非法的bzip2数据错误信息

功能说明：StructuralError其实是一个string，他实现了error接口，用于很方便的返回非法的bzip2数据的错误信息

示例：

	package main
	
	import (
		"archive/tar"
		"compress/bzip2"
		"fmt"
		"io"
		"log"
		"os"
		"reflect"
	)
	
	func main() {
		//打开一个非bzip2压缩文件
		bzip2File, _ := os.Open("main.go")
		defer bzip2File.Close()
	
		//尝试进行解压
		r := bzip2.NewReader(bzip2File)
		//使用tar读取输出文件
		tr := tar.NewReader(r)
		for {
			_, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("the type is ", reflect.TypeOf(err))
				log.Fatalln(err)
			}
		}
	}
