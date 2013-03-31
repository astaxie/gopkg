# func TempFile(dir, prefix string) (f *os.File, err error)

参数列表

- dir 创建临时文件的目录
- prefix 创建临时文件名称的前缀(文件名称组成为前缀+随机数)

返回值：

- f 返回创建的临时文件的指针
- error 返回创建临时文件是否成功

功能说明：

本函数主要是用来在指定的dir目录下创建一个新的、使用prefix为前缀的临时文件，并以读写模式打开该文件并返回os.File指针。 如果dir是空字符串，TempFile使用默认用于临时文件的目录（参见os.TempDir函数）。 如果多个程序调用该函数的话，将会创建不同的临时文件（因此是线程安全的）。调用本函数的程序有责任在不需要临时文件时摧毁它.

代码实例：

	package main

	import "fmt"
	import "io/ioutil"

	func main() {
		f, e := ioutil.TempFile("d:/goTest", "temp")
		defer f.Close()
		fmt.Println(f.Name())
		if e != nil {
			fmt.Println("create tempFile error")
			return
		}
		f.WriteString("hello world!")
		b, e1 := ioutil.ReadFile(f.Name())
		if e1 != nil {
			fmt.Println("read error")
			return
		}
		fmt.Println(string(b))
	}
