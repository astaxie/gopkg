## 结构 IPAddr

	type IPAddr struct {
	    IP IP
	}
	
IP地址结构

====
- func (a *IPAddr) Network() string

返回该ip地址的网络类型名称

====
- func (a *IPAddr) String() string

返回该ip地址下的字符串形式IP

