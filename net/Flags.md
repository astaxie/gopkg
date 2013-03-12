## 结构　Flags

	type Flags uint
	
	const (
	    FlagUp           Flags = 1 << iota // 连接接口标志
	    FlagBroadcast                      // 广播连接接口标志
	    FlagLoopback                       // 环回连接接口标志
	    FlagPointToPoint                   //点对点连接接口标志
	    FlagMulticast                      // 多播连接接口标志
	)


 - func (f Flags) String() string

返回接口标志