# func ReadDir(dirname string) ([]os.FileInfo, error)

参数列表

- dirname 读取目录的目录路径 

返回值：

- []os.FileInfo 返回指定dirname目录下的一个有序的子目录信息列表
- error 返回读取是否成功

功能说明：

本函数主要是用来读取指定dirname目录，并返回一个有序的子目录信息列表

代码实例：

	package main

	import "fmt"
	import "io/ioutil"

	func main() {
		dir_list, e := ioutil.ReadDir("d:/goTest")
		if e != nil {
			fmt.Println("read dir error")
			return
		}
		for i, v := range dir_list {
			fmt.Println(i, "=", v.Name())
		}
	}
