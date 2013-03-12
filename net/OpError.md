## 结构 OpError

	type OpError struct {
	    Op   string
	    Net  string
	    Addr Addr
	    Err  error
	}

请参考Error包

====
- func (*OpError) Error

- func (e *OpError) Temporary() bool

- func (e *OpError) Timeout() bool
