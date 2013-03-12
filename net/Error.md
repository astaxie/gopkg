## 接口 Error

网络错误接口

	type Error interface {
	    error
	    Timeout() bool   // 是否属于一个超时错误
	    Temporary() bool // 是否是一个临时错误
	}
	