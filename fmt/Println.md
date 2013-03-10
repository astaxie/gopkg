# func Println(a ...interface{}) (n int, err error)

参数列表

- a... 值变量列表

返回值：

- 返回打印字符数 n
- 返回error

功能说明：

这个函数主要是用来根据系统默认格式字符串和参数表生成一个打印字符串

代码实例：

 	package main
	
	import 	"fmt"
		
	func main() {
		fmt.Println("默认格式打印!")
	}
