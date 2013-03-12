## func Pipe() (Conn, Conn)

参数列表:

- 无

返回列表:

- Conn 连接接口
- Conn 连接接口

创建一个同步的,全内存的,无缓冲的,全双工管道,并返回2个连接接口

[百度百科 - 管道]()

代码实例:

	package main

	import "fmt"
	import "net"

	func main() {
		conn1, conn2 := net.Pipe()
		fmt.Printf("conn1: %#v \n conn2:%#v", conn1, conn2) 
		//返回类似
		//conn1: &net.pipe{PipeReader:(*io.PipeReader)(0xf840043048), PipeWriter:(*io.PipeWriter)(0xf840043060)} 
	 	//conn2:&net.pipe{PipeReader:(*io.PipeReader)(0xf840043058), PipeWriter:(*io.PipeWriter)(0xf840043050)}
	}
