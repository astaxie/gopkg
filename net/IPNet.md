## 结构 IPNet

	type IPNet struct {
	    IP   IP     // network number
	    Mask IPMask // network mask
	}
	
IPNet 结构

====
- func (n *IPNet) Contains(ip IP) bool

返回此IP网络是否包含该IP

====
- func (n *IPNet) Network() string

返回此IP网络名称  "ip+net"

====
- func (n *IPNet) String() string

返回IP网络的字符串形式