## func Print(v ...interface{})

参数列表：

- v 待输出参数列表

返回值：

- 无

功能说明：

输出日志到标准logger。参数处理方式同fmt.Print

代码实例：

	package main

	import(
		"log"
	)

	func main(){
		log.Print("string", 1, 2.3)//2013/03/10 17:26:06 string1 2.3
	}

