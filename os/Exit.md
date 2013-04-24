## func Exit(code int)

参数列表

- code 程序退出状态码

返回值：

- 无

功能说明：

这个函数主要是立即终止当前程序，defer的函数不会运行

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("running...\n")
        defer fmt.Printf("defer function\n")
        os.Exit(1)
        fmt.Printf("finished\n")
    }

代码输出：

    //test in ArchLinux
    running...
    exit status 1
