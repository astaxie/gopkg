## func Flags() int

参数列表：

- 无

返回值：

- 默认logger的配置值

功能说明：

返回默认logger配置值。

代码实例：

	package main

	import(
		"log"
		"fmt"
	)

	func main(){

		fmt.Println("standard flags :", log.Flags())

		//the flags constants
		fmt.Println(log.Ldate)
		fmt.Println(log.Ltime)
		fmt.Println(log.Lmicroseconds)
		fmt.Println(log.Llongfile)
		fmt.Println(log.Lshortfile)
		fmt.Println(log.LstdFlags)	//LstdFlags     = Ldate | Ltime
	}

