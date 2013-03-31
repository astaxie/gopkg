# func Fprintln(w io.Writer,a ...interface{}) (n int, err error)

参数列表

- w 写入文件指针
- a... 值变量列表

返回值：

- 返回打印字符数 n
- 返回error

功能说明：

这个函数主要是用来根据默认格式字符串和参数表生成一个打印字符串并加换行写入指定文件

代码实例：

 	package main
	
	import 	"fmt"
	import  "os"
		
	func main() {
		fmt.Fprintln(os.Stdout,"默认格式加换行打印!")
	}
