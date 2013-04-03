## func SetOutput(w io.Writer)

参数列表：

- w 目标流，io.Writer类型

返回值：

- 无

功能说明：

设置标准logger的输出目标

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

		log.SetOutput(file)

		log.Println("log to file")
	}

