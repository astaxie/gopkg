## func ServeRequest(codec ServerCodec) error

参数列表
-codec 用使用的ServerCodec

返回值：


功能说明：

这个函数类似ServeCodec,不过这个是同步的。

代码实例：

    package main

    import (
        "bufio"
        "encoding/gob"
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
        // for {
        conn, e := l.Accept()
        if e != nil {
            log.Fatal("Accept error:", e)
        }
        buf := bufio.NewWriter(conn)
        codec := &myServerCodec{conn, gob.NewDecoder(conn), gob.NewEncoder(buf), buf}
        e = rpc.ServeRequest(codec)
        if e != nil {
            log.Fatal("ServeRequst Error", e)
        }
        // }
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

        //创建一个客户端
        client, _ := rpc.Dial("tcp", "127.0.0.1:1234")
        defer client.Close()

        args := &Args{7, 8}
        var reply int
        err := client.Call("Arith.Multiply", args, &reply)
        if err != nil {
            log.Fatal("arith error:", err)
        }
        log.Println(reply)

    }



结果:
    2013/03/20 22:57:00 调用:ReadRequestHeader
    2013/03/20 22:57:00 调用:ReadRequestBody
    2013/03/20 22:57:00 调用:WriteResponse
    2013/03/20 22:57:00 56

