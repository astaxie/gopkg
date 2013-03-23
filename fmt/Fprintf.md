# func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

参数列表

- w 写入文件指针
- format 打印的格式说明 
- a... 值变量列表

返回值：

- 返回打印字符数 n
- 返回error

功能说明：

这个函数主要是用来根据说明格式字符串和参数表生成一个打印字符串

代码实例：

 	package main
	
	import 	"fmt"
	import  "os"
		
	func main() {
		fmt.Fprintf(os.Stdout,"Format:%s\n","格式打印!")
	}
