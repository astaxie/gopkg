## func NewClient(conn io.ReadWriteCloser) *Client

参数列表
-conn 网络连接对象

返回值：
-Client 客户端，返回一个连接到rpc服务的客户端对象

功能说明：

这个函数是用来连接rpc服务，并创建一个客户端对象。 

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

        //暂停2秒，让服务器有足够的时间开启
        time.Sleep(2 * time.Second)

        //设置客户端
        //使用TCP的方式创建一个连接
        conn, err := net.Dial("tcp", "127.0.0.1:1234")
        if err != nil {
            log.Fatal("dial error", err)
        }
        //创建一个客户端对象
        client := rpc.NewClient(conn)
        defer client.Close()

        args := &Args{7, 8}
        var reply int
        //呼叫请求
        err = client.Call("Arith.Multiply", args, &reply)
        if err != nil {
            log.Fatal("arith error:", err)
        }
        log.Println(reply)

    }




结果: 2013/03/21 20:19:30 56

