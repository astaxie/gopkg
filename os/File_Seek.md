## func (f *File) Seek(offset int64, whence int) (ret int64, err error)

参数列表

- offset 文件指针的位置
- whence 相对位置标识

返回值：

- 返回ret 返回文件指针的位置
- 返回err 返回error错误对象

功能说明：

这个函数主要是把文件指针移动到文件的指定位置，whence为0时代表相对文件开始的位置，1代表相对当前位置，2代表相对文件结尾的位置

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        b := make([]byte, 10)
        fi, err := os.Open("Hello.go")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        defer fi.Close()
        fi.Seek(3, 0)
        n, _ := fi.Read(b)
        fmt.Printf("%d\n", n)
        fmt.Printf("%s\n", b[:n])
    }

代码输出：

    //test in ArchLinux
    //Hello.go内容为"aa\nbb\n"
    3
    bb
