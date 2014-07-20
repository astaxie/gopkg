## func ParseMAC(s string) (hw HardwareAddr, err error)

参数列表:

- s 符合 IEEE 802 MAC-48, EUI-48 或 EUI-64 标准的MAC地址字符串。支持以下格式
    - 01:23:45:67:89:ab
    - 01:23:45:67:89:ab:cd:ef
    - 01-23-45-67-89-ab
    - 01-23-45-67-89-ab-cd-ef
    - 0123.4567.89ab
    - 0123.4567.89ab.cdef

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
		fmt.Printf("%#v",hw) 
	}
	
函数输出：

    net.HardwareAddr{0x1, 0x23, 0x45, 0x67, 0x89, 0xab}
