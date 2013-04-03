## 结构 InvalidAddrError

	type InvalidAddrError string

非法地址错误,请查看Error包
	
====
- func (e InvalidAddrError) Error() string

- func (e InvalidAddrError) Temporary() bool

- func (e InvalidAddrError) Timeout() bool

