## func Symlink(oldname, newname string) error

参数列表

- oldname 文件名
- newname 新链接名

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是给一个文件创建一个软链接

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        if err := os.Symlink("Hello.go", "Lhello.go"); err != nil {
            fmt.Printf("Error: %v", err)
            return
        }
        fmt.Printf("The Hello.go's  symbolic link Lhello.go has been created!\n")
    }

代码输出：

    //test in ArchLinux
    The Hello.go's  symbolic link Lhello.go has been created!
