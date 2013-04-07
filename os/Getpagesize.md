## func Getpagesize() int

参数列表

- 无

返回值：

- int 返回系统内在的大小，单位为MB

功能说明：

这个函数主要是返回系统内存的大小

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("%d\n", os.Getpagesize())
    }

代码输出：

    //test in ArchLinux
    4096
