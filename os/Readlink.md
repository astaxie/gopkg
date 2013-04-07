## func Readlink(name string) (string, error)

参数列表

- name 链接名

返回值：

- string 目标文件
- error 返回错误信息对象

功能说明：

这个函数主要是返回一个链接的目标文件

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        link, err := os.Readlink("src")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("%s\n", link)
    }

代码输出：

    //test in ArchLinux
    // src -> /home/miraclesu/go/own/src
    /home/miraclesu/go/own/src
