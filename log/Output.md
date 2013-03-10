## func (l *Logger) Output(calldepth int, s string) error

参数列表：

- calldepth 深度 
- s 字符串

返回值：

- error 错误

功能说明：

输出日志事件。字符串s包含待打印内容，跟在预定义的prefix后面，并且根据flags设置会有区分。如果s末尾没有换行符，这个方法会默认加上一个。calldepth目前预定义均为2,以后会用来支持通用场景，支持其他值配置。（本人注：日志输出不建议直接使用该方法）

代码实例：

	package main

	import(
		"log"
		"os"
	)

	func main(){

		l := log.New(os.Stdout, "log->", log.Ldate)

		l.Output(2, "log output")
	}

