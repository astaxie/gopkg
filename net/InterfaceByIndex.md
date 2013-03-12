## func InterfaceByIndex(index int) (*Interface, error)

参数列表:

- index 索引号

返回列表:

- I*nterface 接口结构指针
- error 错误信息

根据索引查找网络接口

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		interf,err := net.InterfaceByIndex(1)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(interf) //返回 &{1 16436 lo  up|loopback}
	}