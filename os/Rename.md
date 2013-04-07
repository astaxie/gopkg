## func Rename(oldname, newname string) error

参数列表

- oldname 原文件名
- newname 新文件名

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是给一个文件重命名

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        err := os.Rename("Hello.go", "World.go")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("Hello.go has been renamed to World.go!\n")
    }

代码输出：

    //test in ArchLinux
    Hello.go has been renamed to World.go!
