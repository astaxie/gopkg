## func CIDRMask(ones, bits int) IPMask

参数列表:

- ones 网络字符串
- bits 服务名

返回列表:

- IPMask 掩码结构

[百度百科 - CIDR](http://baike.baidu.com/view/4217886.htm)

返回CIDR规范的掩码
ones 必须 大于0 小于 bits
bits 必须为 4(IPv4的字节数)或16(IPv6的字节数)的8倍


代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		mask := net.CIDRMask(2,32)
		fmt.Printf("%#V",mask) //返回 %!V(net.IPMask=[192 0 0 0])
	}
