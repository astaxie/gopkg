## 结构 IP

	type IP []byte
	

ip地址结构

====
- func (ip IP) DefaultMask() IPMask

获取默认IP掩码

====
- func (ip IP) Equal(x IP) bool

判断两个IP是否相等,x 里可以是IPv6

====
- func (ip IP) IsGlobalUnicast() bool

判断是否全球单播地址

[百度百科 - 单播](http://baike.baidu.com/view/625843.htm)

====
- func (ip IP) IsInterfaceLocalMulticast() bool

判断是否本地多播接口

====
- func (ip IP) IsLinkLocalMulticast() bool

判断是否本地多播连接

====
- func (ip IP) IsLinkLocalUnicast() bool

判断是否本地单播连接

====
- func (ip IP) IsLoopback() bool

判断是否环回地址

====
- func (ip IP) IsMulticast() bool

判断是否多播地址

====
- func (ip IP) IsUnspecified() bool

判断ip地址是否是不确定地址

====
- func (ip IP) Mask(mask IPMask) IP

设置网络掩码

====
- func (ip IP) String() string

返回ip地址的字符串形式

====
- func (ip IP) To16() IP

返回16位ip地址

====
- func (ip IP) To4() IP

返回4位IP地址