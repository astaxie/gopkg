## func IPv4Mask(a, b, c, d byte) IPMask

参数列表:

- a a类网络地址
- b b类网络地址
- c c类网络地址
- d d类网络地址

返回列表:

- IPMask 掩码

查找特定网络的端口和服务

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		mask := net.IPv4Mask(byte(127),byte(127),byte(127),byte(127))
		fmt.Printf("%#v",mask) //返回 []byte{0x7f, 0x7f, 0x7f, 0x7f}
	}
