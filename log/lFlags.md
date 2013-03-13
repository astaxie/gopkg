## func (l *Logger) Flags() int

参数列表：

- 无

返回值：

- 当期logger的配置值

功能说明：

返回当前logger配置值。

代码实例：

	package main

	import(
		"log"
		"fmt"
		"os"
	)

	func main(){

		l := log.New(os.Stdout, "", log.LstdFlags)
		fmt.Println("logger l's flags :", l.Flags())
	}


