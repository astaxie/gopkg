## func (server *Server) ServeCodec(codec ServerCodec)

参数列表
-codec 用使用的ServerCodec

返回值：


功能说明：

这个函数是运行一个单连接的DefaultServer，服务一直开启，直到客户端断开连接，并且可以设置ServeCodec，用于处理请求和回复。

代码实例：

    package main

    import (
        "bufio"
        "encoding/gob"
        "errors"
        "io"
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

    type myServerCodec struct {
        rwc    io.ReadWriteCloser
        dec    *gob.Decoder
        enc    *gob.Encoder
        encBuf *bufio.Writer
    }

    func (c *myServerCodec) ReadRequestHeader(r *rpc.Request) error {
        log.Println("调用:ReadRequestHeader")
        return c.dec.Decode(r)
    }

    func (c *myServerCodec) ReadRequestBody(body interface{}) error {
        log.Println("调用:ReadRequestBody")
        return c.dec.Decode(body)
    }

    func (c *myServerCodec) WriteResponse(r *rpc.Response, body interface{}) (err error) {
        log.Println("调用:WriteResponse")
        if err = c.enc.Encode(r); err != nil {
            return
        }
        if err = c.enc.Encode(body); err != nil {
            return
        }
        return c.encBuf.Flush()
    }

    func (c *myServerCodec) Close() error {
        log.Println("调用:Close")
        return c.rwc.Close()
    }

    func myAccept(l net.Listener) {
        //设置服务端
        arith := new(Arith)
        //注册服务
        server := rpc.NewServer()
        server.RegisterName("Serve", arith)
        server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
        conn, e := l.Accept()
        if e != nil {
            log.Fatal("Accept error:", e)
        }
        //创建一个自定义的ServerCode，实际上这些代码就从go源码中获取出来的
        buf := bufio.NewWriter(conn)
        codec := &myServerCodec{conn, gob.NewDecoder(conn), gob.NewEncoder(buf), buf}
        //运行连接
        go server.ServeCodec(codec)
    }
    func main() {

        //创建连接监听
        l, e := net.Listen("tcp", ":1234")
        if e != nil {
            log.Fatal("listen error:", e)
        }
        //运行服务监听
        go myAccept(l)

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
        err = client.Call("Serve.Multiply", args, &reply)
        if err != nil {
            log.Fatal("arith error:", err)
        }
        log.Println(reply)

    }


结果:
    2013/03/21 21:07:59 56

