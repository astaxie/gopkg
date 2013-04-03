## func (server *Server) ServeConn(conn io.ReadWriteCloser)

参数列表
-conn 网络连接对象

返回值：


功能说明：

这个函数是运行一个单连接的DefaultServer，服务一直连接，直到客户端断开连接。

代码实例：

    package main

    import (
        "errors"
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
    func (t *Arith) Divide(args *Args, quo *Quotient) error {
        if args.B == 0 {
            return errors.New("divide by zero")
        }
        quo.Quo = args.A / args.B
        quo.Rem = args.A % args.B
        return nil
    }

    func myAccept(l net.Listener) {
        arith := new(Arith)
        //注册服务
        server := rpc.NewServer()
        server.Register(arith)
        server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
        conn, e := l.Accept()
        if e != nil {
            log.Fatal("Accept error:", e)
        }
        go server.ServeConn(conn)
    }
    func main() {
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
        address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
        if err != nil {
            panic(err)
        }
        //使用TCP的方式进行网络请求
        conn, _ := net.DialTCP("tcp", nil, address)
        defer conn.Close()
        //创建一个客户端
        client := rpc.NewClient(conn)
        defer client.Close()

        args := &Args{7, 8}
        var reply int
        err = client.Call("Arith.Multiply", args, &reply)
        if err != nil {
            log.Fatal("arith error:", err)
        }
        log.Println(reply)

    }



结果:
   2013/03/21 21:38:43 56

