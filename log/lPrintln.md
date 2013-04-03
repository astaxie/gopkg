## func (l *Logger) Println(v ...interface{})

参数列表：

- v 待输出参数列表

返回值：

- 无

功能说明：

调用Output打印日志到当前logger，参数处理方式同fmt.Println

代码实例：

	package main

	import(
		"log"
		"os"
	)

	func main(){

		l := log.New(os.Stdout, "", log.LstdFlags)
		l.Println("hello")	//2013/03/10 17:35:28 hello\n
	}

