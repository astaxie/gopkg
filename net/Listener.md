## 接口 Listener

	type Listener interface {
	    // 等待并返回下一个连接或错误
	    Accept() (c Conn, err error)
	
	    // 关闭这个监听者
	    Close() error
	
	    // 返回监听信息
	    Addr() Addr
	}

监听者接口


	
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
	    log.Fatal(err)
	}
	for {
	    // Wait for a connection. 
	    conn, err := l.Accept()
	    if err != nil {
	        log.Fatal(err)
	    }
	    // Handle the connection in a new goroutine.
	    // The loop then returns to accepting, so that
	    // multiple connections may be served concurrently.
	    go func(c net.Conn) {
	        // Echo all incoming data.
	        io.Copy(c, c)
	        // Shut down the connection.
	        c.Close()
	    }(conn)
	}