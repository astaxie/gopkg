## func MkdirAll(path string, perm FileMode) error

参数列表

- path 目录路径
- perm 权限

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是按文件目录路径新建目录，如果路径中的上一级目录不存在，会自动创建

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        if err := os.MkdirAll("hello_go/world", 0777); err != nil {
            fmt.Printf("Error: %v\n", err)
        }
        fmt.Printf("hello_go/world has been created!\n")
    }

代码输出：

    //test in ArchLinux
    hello_go/world has been created!
