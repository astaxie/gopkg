## func IPv4Mask(a, b, c, d byte) IPMask

参数列表:

- a IP掩码的第1个字节
- b IP掩码的第2个字节
- c IP掩码的第3个字节
- d IP掩码的第4个字节

返回列表:

- IPMask 掩码

函数功能

- 生成IP掩码

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		mask := net.IPv4Mask(127, 127, 127, 127)
		fmt.Printf("%#v",mask)
	}

代码输出：

    net.IPMask{0x7f, 0x7f, 0x7f, 0x7f}