## func (server *Server) HandleHTTP(rpcPath, debugPath string)

参数列表
-rpcPath rpc路径
-debugPath debug路径

返回值：



功能说明：

这个函数是将rpc的服务注册到http协议上，使用http的方式进行接受数据和传送数据，并可以设置 rpcPath和debugPath。

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
        server := rpc.NewServer()
        server.Register(arith)
        server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
        //创建连接监听
        l, e := net.Listen("tcp", ":1234")
        if e != nil {
            log.Fatal("listen error:", e)
        }
        //设置连接监听
        go server.Accept(l)

        //暂停2秒，让服务器有足够的时间开启
        time.Sleep(2 * time.Second)

        //使用TCP的方式进行网络请求
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



结果:2013/03/21 21:01:02 56

