## func (f *File) Truncate(size int64) error

参数列表

- size 截断后的大小

返回值：

- 返回error 返回error错误对象

功能说明：

这个函数主要是把文件截断到指定大小

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fi, err := os.OpenFile("Hello.go", os.O_RDWR, 0420)
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        defer fi.Close()
        f, _ := fi.Stat()
        fmt.Printf("file size: %+v\n", f.Size())
        fi.Truncate(6)
        f, _ = fi.Stat()
        fmt.Printf("now file size: %+v\n", f.Size())
    }

代码输出：

    //test in ArchLinux
    file size: 9
    now file size: 6
