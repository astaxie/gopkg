## func Panicln(v ...interface{})

参数列表：

- v 待输出参数列表

返回值：

- 无

功能说明：

相当于调用Println()并调用panic()

代码实例：

	package main

	import(
		"log"
	)

	func main(){

		defer func(){

			if err := recover(); err != nil{
				if err == "3q\n"{
					log.Println("you are welcome")
				}
			}
		}()

		log.Panicln("3q")
	}

	//this will print like this:
	//2013/03/10 16:48:06 3q
	//2013/03/10 16:48:06 you are welcome

