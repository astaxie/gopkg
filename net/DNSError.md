## 结构 DNSError

DNSError 错误结构 - 请参考 Error包

结构:

	type DNSError struct {
    		Err       string // 错误说明
    		Name      string // 错误名
    		Server    string // 使用的服务名
    		IsTimeout bool // 是否超时
	}


- func (*DNSError) Error

- func (*DNSError) Temporary

- func (*DNSError) Timeout

