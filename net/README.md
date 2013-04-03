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
- [LookupPort](LookupPort.md)
- [LookupSRV](LookupSRV.md)
- [LookupTXT](LookupTXT.md)
- [SplitHostPort](SplitHostPort.md)
- [Dial](Dial.md)
- [DialTimeout](DialTimeout.md)
- [Pipe](Pipe.md)
- [ParseMAC](ParseMAC.md)
- [IPv4](IPv4.md)
- [ParseCIDR](ParseCIDR.md)
- [ParseIP](ParseIP.md)
- [ResolveIPAddr](ResolveIPAddr.md)
- [DialIP](DialIP.md)
- [ListenIP](ListenIP.md)
- [CIDRMask](CIDRMask.md)
- [IPv4Mask](IPv4Mask.md)
- [InterfaceByIndex](InterfaceByIndex.md)
- [InterfaceByName](InterfaceByName.md)
- [FileListener](FileListener.md)
- [Listen](Listen.md)

## 结构

- [Addr](Addr.md)

- [AddrError](AddrError.md)
 - [Error](AddrError.md)
 - [Temporary](AddrError.md)
 - [Timeout](AddrError.md)

- [Conn](Conn.md)
 - [Read](Conn.md)
 - [Write](Conn.md)
 - [Close](Conn.md)
 - [LocalAddr](Conn.md)
 - [RemoteAddr](Conn.md)
 - [SetDeadline](Conn.md)
 - [SetReadDeadline](Conn.md)
 - [SetWriteDeadline](Conn.md)

- [DNSConfigError](DNSConfigError.md)
 - [Error](DNSConfigError.md)
 - [Temporary](DNSConfigError.md)
 - [Timeout](DNSConfigError.md)

- [DNSError](DNSError.md)
 - [Error](DNSError.md)
 - [Temporary](DNSError.md)
 - [Timeout](DNSError.md)

- [Error](Error.md)

- [Flags](Flags.md)
 - [String](Flags.md)

- [HardwareAddr](HardwareAddr.md)
 - [String](HardwareAddr.md)

- [IP](IP.md)
 - [DefaultMask](IP.md)
 - [Equal](IP.md)
 - [IsGlobalUnicast](IP.md)
 - [IsInterfaceLocalMulticast](IP.md)
 - [IsLinkLocalMulticast](IP.md)
 - [IsLinkLocalUnicast](IP.md)
 - [IsLoopback](IP.md)
 - [IsMulticast](IP.md)
 - [IsUnspecified](IP.md)
 - [Mask](IP.md)
 - [String](IP.md)
 - [To16](IP.md)

- [IPAddr](IPAddr.md)
 - [Network](IPAddr.md)
 - [String](IPAddr.md)

- [IPConn](IPConn.md)
 - [Close](IPConn.md)
 - [File](IPConn.md)
 - [LocalAddr](IPConn.md)
 - [Read](IPConn.md)
 - [ReadFrom](IPConn.md)
 - [ReadFromIP](IPConn.md)
 - [RemoteAddr](IPConn.md)
 - [SetDeadline](IPConn.md)
 - [SetReadBuffer](IPConn.md)
 - [SetReadDeadline](IPConn.md)
 - [SetWriteBuffer](IPConn.md)
 - [SetWriteDeadline](IPConn.md)
 - [Write](IPConn.md)
 - [WriteTo](IPConn.md)
 - [WriteToIP](IPConn.md)

- [IPMask](IPMask.md)
 - [Size](IPMask.md)
 - [String](IPMask.md)
	
- [IPNet](IPNet.md)
 - [Contains](IPNet.md)
 - [Network](IPNet.md)
 - [String](IPNet.md)

- [Interface](Interface.md)
 - [Addrs](Interface.md)
 - [MulticastAddrs](Interface.md)

- [InvalidAddrError](InvalidAddrError.md)
 - [Error](InvalidAddrError.md)
 - [Temporary](InvalidAddrError.md)
 - [Timeout](InvalidAddrError.md)

- [Listener](Listener.md)

- [MX](MX.md)

- [OpError](OpError.md)

