## func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error

参数列表
-serviceMethod 服务名称
-args 发送的参数
-reply 回复
返回值：
-error 错误信息

功能说明：

这个函数是用来调用rpc服务。

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
        client, _ := rpc.Dial("tcp", "127.0.0.1:1234")

        args := &Args{7, 8}
        var reply int
        err := client.Call("Arith.Multiply", args, &reply)
        if err != nil {
            log.Fatal("arith error:", err)
        }
        log.Println(reply)

    }





结果:2013/03/21 20:29:45 56

