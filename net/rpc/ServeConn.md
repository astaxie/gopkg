## func ServeConn(conn io.ReadWriteCloser)
参数列表
-conn 网络连接对象

返回值：


功能说明：

这个函数是运行一个单连接的DefaultServer，服务一直连接，直到客户端断开连接。实际上就是调用DefaultServer.ServeConn(conn)这个方法。

代码实例：

    package main

    import (
        "log"
        "net"
        "net/rpc"
        "time"
    )

    type Args struct {
        A, B int
    }
    type Quotient struct {
        Quo, Rem int
    }
    type Arith int

    func (t *Arith) Multiply(args *Args, reply *int) error {
        *reply = args.A * args.B
        return nil
    }

    func myAccept(l net.Listener) {
        conn, e := l.Accept()
        if e != nil {
            log.Fatal("Accept error:", e)
        }
        go rpc.ServeConn(conn)
    }
    func main() {
        //设置服务端
        arith := new(Arith)
        //注册服务
        rpc.Register(arith)
        //将服务注册到了HTTP协议
        rpc.HandleHTTP()
        //创建连接监听
        l, e := net.Listen("tcp", ":1234")
        if e != nil {
            log.Fatal("listen error:", e)
        }
        //设置连接监听
        go myAccept(l)
        //设置自定义的serverCodec

        //暂停2秒，让服务器有足够的时间开启
        time.Sleep(2 * time.Second)

        //设置客户端
        client, _ := rpc.Dial("tcp", "127.0.0.1:1234")
        defer client.Close()

        args := &Args{7, 8}
        var reply int
        //呼叫请求
        err := client.Call("Arith.Multiply", args, &reply)
        if err != nil {
            log.Fatal("arith error:", err)
        }
        log.Println(reply)

    }



结果:
    2013/03/20 22:49:30 56

