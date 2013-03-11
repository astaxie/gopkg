## func Notify(c chan<- os.Signal, sig ...os.Signal)

参数列表
- c 信号 channel
- sig 需要捕获的信号

实例：
    package main

    import(
        "fmt"
        "os"
        "os/signal"
    )

    func main() {
        schan := make(chan os.Signal)
        go signal.Notify(schan, os.Interrupt, os.Kill)
        fmt.Printf("Signal captured: %d", <-schan)
    }
