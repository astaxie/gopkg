## func Prefix() string

参数列表：

- 无

返回值：

- 标准logger前缀，字符串类型

功能说明：

返回标准logger的输出前缀

代码实例：

	package main

	import(
		"log"
		"fmt"
	)

	func main(){
	
		fmt.Print(log.Prefix())	//this will print nothing

		log.Println(1)	//2013/03/10 17:02:05 1

		log.SetPrefix("log:")
		fmt.Println(log.Prefix())	//log:

		log.Println(2)	//log:2013/03/10 17:02:05 2
	}

