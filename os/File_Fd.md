## func (f *File) Fd() uintptr

参数列表

- 无

返回值：

- 返回unitptr 返回文件句柄

功能说明：

这个函数主要是获取文件句柄

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fi, err := os.Open("Hello.go")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        defer fi.Close()
        fmt.Printf("%v\n", fi.Fd())
    }

代码输出：

    //test in ArchLinux
    3
