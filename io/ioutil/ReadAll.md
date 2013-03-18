# func ReadAll(r io.Reader) ([]byte, error)

参数列表

- r 读取对象 

返回值：

- []byte 返回读取对象的内容
- error 返回读取是否成功

功能说明：

本函数主要是用来从r读取直到遇到error或EOF并返回读取的数据；成功的读取返回的err为nil，而不是EOF。因为ReadAll定义为从资源读取数据直到EOF，它不会将从r读取的EOF视为应该报告的错误

代码实例：

	package main

	import "fmt"
	import "io/ioutil"
	import "strings"

	func main() {
		s := strings.NewReader("hello world!")
		b, e := ioutil.ReadAll(s)
		if e != nil {
			fmt.Printf("%v\n", e)
			return
		}
		fmt.Println(string(b))
	}
