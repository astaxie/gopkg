## 类型 Addr

Addr是一个接口,代表一个网络端的地址

结构:

	type Addr interface {
   	 	Network() string // 网络名字符串
   	 	String() string  // 地址字符串
	}
