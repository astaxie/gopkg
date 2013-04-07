## func Link(oldname, newname string) error

参数列表

- oldname 源文件名
- newname 目标文件名

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是给一个文件建立一个硬链接

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        if err := os.Link("Hello.go", "Lhello.go"); err != nil {
            fmt.Printf("Error: %v\n", err)
        }
        fmt.Printf("Lhello.go has created!\n")
    }

代码输出：

    //test in ArchLinux
    Lhello.go has been created!
