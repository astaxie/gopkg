## func Getpid() int

参数列表

- 无

返回值：

- int 返回调用者进程id

功能说明：

这个函数主要是返回调用者进程id

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("%d\n", os.Getpid())
    }

代码输出：

    //test in ArchLinux
    //每次执行的结果可能都不一样
    11396
