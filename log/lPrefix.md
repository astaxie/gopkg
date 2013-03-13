## func (l *Logger) Prefix() string

参数列表：

- 无

返回值：

- logger前缀，字符串类型

功能说明：

返回当前logger的输出前缀

代码实例：

	package main

	import(
		"log"
		"fmt"
		"os"
	)

	func main(){

		l := log.New(os.Stdout, "", log.LstdFlags)
		fmt.Print(l.Prefix())	//this will print nothing

		l.Println(1)	//2013/03/10 17:02:05 1

		l.SetPrefix("log:")
		fmt.Println(l.Prefix())	//log:

		l.Println(2)	//log:2013/03/10 17:02:05 2
	}

