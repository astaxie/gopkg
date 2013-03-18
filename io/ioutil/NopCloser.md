# func NopCloser(r io.Reader) io.ReadCloser

参数列表

- r 读取对象 

返回值：

- 返回一个读取对象ReadCloser接口，该接口仅提供Close方法。

功能说明：

本函数主要是用来返回一个读取对象的ReadCloser接口，该接口仅提供Close方法

代码实例：

	package main

	import "io/ioutil"
	import "strings"

	func main() {
		s := strings.NewReader("hello world!")
		r := ioutil.NopCloser(s)
		defer r.Close()
	}
