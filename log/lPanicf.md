## func (l *Logger) Panicf(format string, v ...interface{})

参数列表：

- format 	输出格式
- v			待输出参数列表

返回值：

- 无

功能说明：

相当于调用l.Printf()，之后调用panic()

代码实例：

	package main

	import(
		"log"
		"os"
	)

	func main(){

		defer func(){

			if err := recover(); err != nil{
				if err == "3q"{
					log.Println("you are welcome")
				}
			}
		}()

		l := log.New(os.Stdout, "", log.LstdFlags)
		l.Panicf("%d%s", 3, "q")
	}

	//this will print like this:
	//2013/03/10 16:48:06 3q
	//2013/03/10 16:48:06 you are welcome


