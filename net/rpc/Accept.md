## func Accept(lis net.Listener)

参数列表

-lis 要设置的网络监听对象

返回值：



功能说明：

这个函数主要来设置网络连接监听，以响应每次的网络连接请求。如果看源码就会发现，这个方法实际上是调用DefaultServer.Accept(lis)这个方法。这个方法一般用go rpc.Accept(l)进行调用。 

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

    func main() {
        //设置服务端
        arith := new(Arith)
        //注册服务
        rpc.Register(arith)
        //创建连接监听
        l, e := net.Listen("tcp", ":1234")
        if e != nil {
            log.Fatal("listen error:", e)
        }
        //设置连接监听
        go rpc.Accept(l)

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



结果:2013/03/19 20:35:05 56

