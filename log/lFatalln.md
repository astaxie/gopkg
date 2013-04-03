## func (l *Logger) Fatalln(v ...interface{})

参数列表：

- v

返回值：

- 无

功能说明：

打印一行日志并退出。相当于调用l.Println()并os.Exit(1)

代码实例：

	package main

	import(
		"log"
		"os"
	)

	func main(){

		l := log.New(os.Stdout, "", log.LstdFlags)
		l.Fatalln("bye!")//2013/03/10 16:14:54 bye!\n
	}
