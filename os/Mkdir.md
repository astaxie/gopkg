## func Mkdir(name string, perm FileMode) error

参数列表

- name 目录名
- perm 权限

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是新建一个指定权限的目录

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        if err := os.Mkdir("hello_go", 0777); err != nil {
            fmt.Printf("Error: %v\n", err)
        }
        fmt.Printf("hello_go has been created!\n")
    }

代码输出：

    //test in ArchLinux
    hello_go has been created!
