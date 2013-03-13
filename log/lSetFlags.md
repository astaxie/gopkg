## func (l *Logger) SetFlags(flag int)

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
		"os"
	)

	func main(){

		l := log.New(os.Stdout, "", log.LstdFlags)

		l.Println(l.Flags())	//2013/03/10 17:46:53 3

		l.SetFlags(log.Ldate)

		l.Println(l.Flags())	//2013/03/10 1

		l.SetFlags(log.LstdFlags | log.Lshortfile)

		l.Println(l.Flags())	//2013/03/10 17:46:53 setflags.go:17: 19
	}

