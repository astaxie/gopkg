## func ParseMAC(s string) (hw HardwareAddr, err error)

参数列表:

- s MAC字符串

返回列表:

- hw 物理硬件地址
- err 错误信息

将mac地址解析为物理硬件地址结构

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		hw,err := net.ParseMAC("01:23:45:67:89:ab")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%#v",hw) //返回 []byte{0x1, 0x23, 0x45, 0x67, 0x89, 0xab}
	}
