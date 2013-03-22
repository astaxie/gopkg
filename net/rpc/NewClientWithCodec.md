## func NewClientWithCodec(codec ClientCodec) *Client

参数列表
-codec ClientCodec,用户处理客户端网络连接的请求和回复。

返回值：
-Client 客户端，返回一个连接到rpc服务的客户端对象

功能说明：

这个函数是用来连接rpc服务，并返回一个客户端对象，可以设置一个 ClientCodec，来处理请求和回复。

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

    type myClientCodec struct {
        rwc    io.ReadWriteCloser
        dec    *gob.Decoder
        enc    *gob.Encoder
        encBuf *bufio.Writer
    }

    func (c *myClientCodec) ReadResponseHeader(r *rpc.Response) error {
        log.Println("调用:ReadResponseHeader")
        return c.dec.Decode(r)
    }

    func (c *myClientCodec) ReadResponseBody(body interface{}) error {
        log.Println("调用:ReadResponseBody")
        return c.dec.Decode(body)
    }

    func (c *myClientCodec) WriteRequest(r *rpc.Request, body interface{}) (err error) {
        log.Println("调用:WriteRequest")
        if err = c.enc.Encode(r); err != nil {
            return
        }
        if err = c.enc.Encode(body); err != nil {
            return
        }
        return c.encBuf.Flush()
    }

    func (c *myClientCodec) Close() error {
        log.Println("调用:Close")
        return c.rwc.Close()
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
        go rpc.Accept(l)
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
        buf := bufio.NewWriter(conn)
        codec := &myClientCodec{conn, gob.NewDecoder(conn), gob.NewEncoder(buf), buf}
        client := rpc.NewClientWithCodec(codec)
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
    2013/03/22 20:56:35 调用:WriteRequest
    2013/03/22 20:56:35 调用:ReadResponseHeader
    2013/03/22 20:56:35 调用:ReadResponseBody
    2013/03/22 20:56:35 调用:ReadResponseHeader
    2013/03/22 20:56:35 56
    2013/03/22 20:56:35 调用:Close

