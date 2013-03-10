## func New(out io.Writer, prefix string, flag int) *Logger

参数列表：

- out 输出目标
- prefix 输出前缀
- flag 格式配置标识值

返回值：

- 自定义的logger

功能说明：

这个方法用来自定义logger，指定输出目标、格式等

代码实例：

package main

import(
	"log"
	"os"
)

func main(){

	file, err := os.OpenFile("sample.txt", os.O_WRONLY, 0666)
	if err != nil{
		panic(err)
	}

	defer file.Close()

	l := log.New(file, "logger", log.Ldate)

	l.Println("log to file sample")	//logger2013/03/10 log to file sample
}

