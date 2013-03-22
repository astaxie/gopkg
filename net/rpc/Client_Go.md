## func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call

参数列表
-serviceMethod 服务名称
-args 发送的参数
-reply 回复
-done 一个Call类型的channel
返回值：
-Call 返回一个Call

功能说明：

这个函数是用来异步调用rpc服务

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
        //设置自定义的serverCodec

        //暂停2秒，让服务器有足够的时间开启
        time.Sleep(2 * time.Second)

        //使用TCP的方式进行网络请求
        client, _ := rpc.Dial("tcp", "127.0.0.1:1234")

        defer client.Close()

        args := &Args{7, 8}
        var reply int
        call := client.Go("Arith.Multiply", args, &reply, make(chan *rpc.Call, 1))
        call = <-call.Done
        if call.Error != nil {
            log.Println("error:", call.Error)
        }
        log.Println(reply)

    }






结果:2013/03/21 20:56:39 56

