# 结构 UDPAddr

    type UDPAddr struct {
	    IP   IP		// Ip地址
	    Port int	// 端口
	    Zone string // IPv6 寻址空间
    }

UDPAddr代表一个UDP地址

## func ResolveUDPAddr(net, addr string) (*UDPAddr, error)

参数列表

- net UDP类型,可接受字符串:udp, udp4, udp6
- addr 地址的字符串: [host]:[port]

返回值：

- *UDPAddr UDP地址
- error 错误(如果有)

功能说明：

该函数将字符串转换成对应的udp地址,以用于后续操作

代码实例：
    
    package main
    
    import (
	    "fmt"
	    "net"
    )
    
    func main() {
	    addr, err := net.ResolveUDPAddr("udp", ":53")
	    if err != nil {
	    	fmt.Println(err)
	    	return
	    }
	    fmt.Println(addr)
    }
    

## func (a *UDPAddr) Network() string

返回值：

- string 地址类型

功能说明：

该函数返回UDPAddr实例的地址类型

    package main
    
    import (
    	"fmt"
    	"net"
    )
    
    func main() {
    	addr, err := net.ResolveUDPAddr("udp", ":53")
    	if err != nil {
    		fmt.Println(err)
    		return
    	}
    	fmt.Println(addr.Network())
    }

## func (a *UDPAddr) String() string

返回值:

- string 地址的字符串表示

功能说明：

该函数返回UDPAddr实例的字符串地址
    
    package main
    
    import (
    	"fmt"
    	"net"
    )
    
    func main() {
    	addr, err := net.ResolveUDPAddr("udp", ":53")
    	if err != nil {
    		fmt.Println(err)
    		return
    	}
    	fmt.Println(addr.String())
    }
