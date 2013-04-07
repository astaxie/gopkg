## func TempDir() string

参数列表

- 无

返回值：

- 返回string 系统的tmp目录

功能说明：

这个函数主要是获取系统的tmp目录

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("The temp dir is: %s\n", os.TempDir())
    }

代码输出：

    //test in ArchLinux
    The temp dir is: /tmp
