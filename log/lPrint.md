## func (l *Logger) Print(v ...interface{})

参数列表：

- v 待输出参数列表

返回值：

- 无

功能说明：

输出日志到logger。参数处理方式同fmt.Print

代码实例：

	package main

	import(
		"log"
		"os"
	)

	func main(){
		l := log.New(os.Stdout, "", log.LstdFlags)
		l.Print("string", 1, 2.3)//2013/03/10 17:26:06 string1 2.3
	}

