## func (l *Logger) Panic(v ...interface{})

参数列表：

- v 待输出参数列表

返回值：

- 无

功能说明：

这个方法相当于调用l.Print()及panic()

代码实例：

	package main

	import(
		"log"
		"fmt"
		"os"
	)

	func main(){


		l := log.New(os.Stdout, "", log.LstdFlags)

		defer func(){
			if err := recover(); err !=nil{

				fmt.Println(err)	//output : "call panic and stop"
				handleException()
			}
		}()


		l.Panic("call panic and stop")

		log.Println("this will not be called.")
	}

	func handleException(){
		log.Println("recovering...")
	}

