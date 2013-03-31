## func SetFlags(flag int)

参数列表：

- flag logger配置值

返回值：

- 无

功能说明：

这个方法用来设置标准logger的配置，默认为3（logger.LstdFlags）

代码实例：

	package main

	import(
		"log"
	)

	func main(){

		log.Println(log.Flags())	//2013/03/10 17:46:53 3

		log.SetFlags(log.Ldate)

		log.Println(log.Flags())	//2013/03/10 1

		log.SetFlags(log.LstdFlags | log.Lshortfile)

		log.Println(log.Flags())	//2013/03/10 17:46:53 setflags.go:17: 19
	}

