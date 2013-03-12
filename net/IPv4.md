## func IPv4(a, b, c, d byte) IP

参数列表:

- a a类地址
- b b类地址
- c c类地址
- d d类地址

返回列表:

- IP ip结构


根据 a,b,c,d类地址返回ip结构

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		ip := net.IPv4(byte(127),byte(0),byte(1),byte(1))
		fmt.Println(ip) //返回 127.0.1.1
	}