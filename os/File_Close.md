## func (f *File) Close() error

参数列表

- 无

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是关闭文件指针

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
        }
        if err = fi.Close(); err != nil {
            fmt.Printf("Error: %s\n", err)
        }
        fmt.Printf("fi has been closed!")
    }

代码输出：

    //test in ArchLinux
    fi has been closed!
