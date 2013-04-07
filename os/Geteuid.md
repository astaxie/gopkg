## func Geteuid() int

参数列表

- 无

返回值：

- int 调用者的有效用户id

功能说明：

这个函数主要是返回调用者的有效用户id

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("%d\n", os.Geteuid())
    }

代码输出：

    //test in ArchLinux
    1000
