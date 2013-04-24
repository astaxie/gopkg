## func Getgid() int

参数列表

- 无

返回值：

- int 调用者的用户组id

功能说明：

这个函数主要是返回调用者的用户组id

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("%d\n", os.Getgid())
    }

代码输出：

    //test in ArchLinux
    100
