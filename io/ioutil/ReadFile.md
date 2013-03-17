# func ReadFile(filename string) ([]byte, error)

参数列表

- filename 读取文件的名称(包括文件目录及名称) 

返回值：

- []byte 返回读取指定文件的内容
- error 返回读取文件是否成功

功能说明：

本函数主要是用来从指定的filename文件中读取数据并返回文件的内容；成功的调用返回的err为nil，而不是EOF。因为ReadFile定义为从资源读取数据直到EOF，它不会将从r读取的EOF视为应该报告的错误

代码实例：

	package main

	import "fmt"
	import "io/ioutil"

	func main() {
		b, e := ioutil.ReadFile("d:/goTest/123.txt")
		if e != nil {
			fmt.Println("read file error")
			return
		}
		fmt.Println(string(b))
	}
