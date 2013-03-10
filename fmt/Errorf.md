# func Errorf(format string, a ...interface{}) error

参数列表

- format 打印的格式说明 
- a... 值变量列表

返回值：

- 返回error

功能说明：

这个函数主要是用来根据格式字符串和参数表生成一个字符串，该字符串满足error接口

代码实例：

  	package main
	
	import "fmt"
	
	func main() {
		fmt.Errorf("Error:%s\n","Test error!")
	}
