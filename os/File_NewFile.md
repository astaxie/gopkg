## func NewFile(fd uintptr, name string) *File

参数列表

- fd 文件句柄
- name 文件名

返回值：

- file 文件指针

功能说明：

这个函数主要是新建一个文件（但不保存），返回文件指针

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fi, err := os.Open("Hello.go")
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        defer fi.Close()
        file := os.NewFile(fi.Fd(), "World.go")
        defer file.Close()
        fmt.Printf("The %s has been new!\n", file.Name())
    }

代码输出：

    //test in ArchLinux
    The World.go has been new!
