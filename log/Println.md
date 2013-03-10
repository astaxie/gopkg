## func Println(v ...interface{})

参数列表：

- v

返回值：

- 无

功能说明：

调用Output打印日志到标准logger，参数处理方式同fmt.Println

代码实例：

	package main

	import(
		"log"
	)

	func main(){
		log.Println("hello")	//2013/03/10 17:35:28 hello\n
	}

