## 结构 Conn

Conn 是一个接口(interface)

需要实现以下方法:

====
- Read(b []byte) (n int, err error) 
	
	从缓冲区读取数据,当读取超时时可以返回Timeout错误(需要设置超时时间)

====
- Write(b []byte) (n int, err error) 
	
	 写入数据到缓冲区,当写入超时时可以返回Timeout错误(需要设置超时时间)

====
- Close() error

	关闭连接
	
====
- LocalAddr() Addr

	返回本地网络地址
	
====
- RemoteAddr() Addr

	返回远端网络地址
	
====
- SetDeadline(t time.Time) error

	设置读写超时时间
	
====
- SetReadDeadline(t time.Time) error

	设置读取超时时间
	
====
- SetWriteDeadline(t time.Time) error

	设置写入超时时间
		
