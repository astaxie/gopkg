## func Fatal(v ...interface{})

参数列表：

- v 待输出参数列表

返回值：

- 无

功能说明：

打印日志并退出。相当于调用Print()并os.Exit(1)

代码实例：

	package main

	import(
		"log"
	)

	func main(){

		age := 25
		log.Fatal("Hi & Bye ! Age = ", age)// this will print "Hi & Bye ! Age = 25"	

		log.Println("This will not be called.")
	}

