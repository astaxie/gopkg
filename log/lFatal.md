## func (l *Logger) Fatal(v ...interface{})

参数列表：

- v 待输出参数列表

返回值：

- 无

功能说明：

打印日志并退出。相当于调用l.Print()并os.Exit(1)

代码实例：

	package main

	import(
		"log"
		"os"
	)

	func main(){

		l := log.New(os.Stdout, "", log.LstdFlags)
		age := 25
		l.Fatal("Hi & Bye ! Age = ", age)// this will print "Hi & Bye ! Age = 25"

		l.Println("This will not be called.")
	}

