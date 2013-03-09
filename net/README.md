# net包
 
## 概述
net 包是处理网络I/O的包

## 常量

P地址长度

	const (
	    IPv4len = 4
	    IPv6len = 16
	)

## 变量

知名IPv4地址

	var (
    	IPv4bcast     = IPv4(255, 255, 255, 255) // 广播
    	IPv4allsys    = IPv4(224, 0, 0, 1)       // 所有主机的地址 （包括所有路由器地址）
	    IPv4allrouter = IPv4(224, 0, 0, 2)       // 所有组播路由器的地址
	    IPv4zero      = IPv4(0, 0, 0, 0)         // 不确定地址
	)
	

知名IPv6地址

	var (
	    IPv6zero                   = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	    IPv6unspecified            = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	    IPv6loopback               = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	    IPv6interfacelocalallnodes = IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	    IPv6linklocalallnodes      = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	    IPv6linklocalallrouters    = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
	)
	
错误

	var ErrWriteToConnected = errors.New("use of WriteTo with pre-connected UDP")
	

## 函数

- [InterfaceAddrs](InterfaceAddrs.md)
- [Interfaces](Interfaces.md)
- [JoinHostPort](JoinHostPort.md)
- [LookupAddr](LookupAddr.md)
- [LookupCNAME](LookupCNAME.md)
- [LookupHost](LookupHost.md)
- [LookupIP](LookupIP.md)


## 结构

- [Addr](Addr.md)