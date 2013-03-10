## func Printf(format string, v ...interface{})

参数列表：

- format 输出格式
- v 待输出参数列表

返回值：

- 无

功能说明：

调用Output输出日志到标准logger。参数处理方式同fmt.Printf

代码实例：

	package main

	import(
		"log"
	)

	func main(){

		log.Printf("%s", "hello")	//hello

		name := "golang"
		log.Printf("%8d,%8s", 23, name) //2013/03/10 16:08:49       23,  golang
	}

