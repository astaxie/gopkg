## func Getwd() (pwd string, err error)

参数列表

- 无

返回值：

- pwd 当前目录
- err 返回错误信息对象

功能说明：

这个函数主要是返回当前目录

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        pwd, err := os.Getwd()
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        fmt.Printf("The current directory is: %s\n", pwd)
    }

代码输出：

    //test in ArchLinux
    The current directory is: /home/miraclesu/Programs/go
