## func (f *File) Name() string

参数列表

- 无

返回值：

- 返回string 返回文件名

功能说明：

这个函数主要是获取文件名

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
        fmt.Printf("The file name is: %v\n", fi.Name())
    }

代码输出：

    //test in ArchLinux
    The file name is: Hello.go
