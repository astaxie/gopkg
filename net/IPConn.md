## 结构 IPConn

	type IPConn struct {
	    // 空结构
	}
	
IPConn是 Conn接口的实现

====
- func (*IPConn) Close

关闭IP连接

====
- func (c *IPConn) File() (f *os.File, err error)

返回一个独占模式的系统文件结构指针,注意.当该连接关闭或者文件关闭时互不影响

====
- func (c *IPConn) LocalAddr() Addr

返回连接的本地地址

====
- func (c *IPConn) Read(b []byte) (int, error)

读取数据到缓冲区(b),返回读取字节数或者错误

====
- func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)

读取数据到缓冲区(b),返回读取字节数,地址结构,或者错误

====
- func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)

读取数据到缓冲区(b),返回读取字节数,IP地址结构,或者错误

====
- func (c *IPConn) RemoteAddr() Addr

返回远端地址结构

====
- func (c *IPConn) SetDeadline(t time.Time) error

设置读/写超时时间

====
- func (c *IPConn) SetReadBuffer(bytes int) error

设置读缓冲区长度

====
- func (c *IPConn) SetReadDeadline(t time.Time) error

设置读取超时时间

====
- func (c *IPConn) SetWriteBuffer(bytes int) error

设置写缓存区长度

====
- func (c *IPConn) SetWriteDeadline(t time.Time) error

设置写入超时时间

====
- func (c *IPConn) Write(b []byte) (int, error)

往缓冲区写入数据,返回写入缓冲区的字节数或者错误

====
- func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)

往对应地址结构的缓冲区写入数据 ,返回写入缓冲区的字节数或者错误- 实现PacketConn接口方法

====
- func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)

往对应IP地址结构的缓冲区写入数据 ,返回写入缓冲区的字节数或者错误


