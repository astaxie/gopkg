# func TempDir(dir, prefix string) (name string, err error)

参数列表

- dir 创建临时目录的父目录
- prefix 创建临时目录的前缀

返回值：

- name 返回创建的临时目录名称(名称的组成为前缀+随机数)
- error 返回创建临时目录是否成功

功能说明：

本函数主要是用来在指定的dir目录里创建一个新的、使用prfix作为前缀的临时文件夹，并返回文件夹的路径。 如果dir是空字符串，TempDir使用默认用于临时文件的目录（参见os.TempDir函数）。 如果多个程序调用该函数的话，将会创建不同的临时目录（因此是线程安全的）。调用本函数的程序有责任在不需要临时文件夹时摧毁它。

代码实例：

	package main

	import "fmt"
	import "io/ioutil"

	func main() {
		f, e := ioutil.TempDir("d:/goTest", "temp")
		if e != nil {
			fmt.Println("create tempDir error")
			return
		}
		fmt.Println(f)
	}
