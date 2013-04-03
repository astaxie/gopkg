## 结构 Interface

	type Interface struct {
	    Index        int          // 应该是一个正整数,如果为0的话则代表未启用
	    MTU          int          // 最大传输单位
	    Name         string       // 比如 "en0", "lo0", "eth0.100"
	    HardwareAddr HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form
	    Flags        Flags        // 比如 FlagUp, FlagLoopback, FlagMulticast
	}

网络接口结构(译者注:你可以暂时理解为网卡/网关/路由信息)

====
- func (ifi *Interface) Addrs() ([]Addr, error)

获取该网络接口下所有地址

====
- func (ifi *Interface) MulticastAddrs() ([]Addr, error)

获取该网络接口下所有多播地址

