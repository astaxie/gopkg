## func (f *File) Write(b []byte) (n int, err error)

参数列表

- b 要写入的内容

返回值：

- 返回int 返回写入的字节数
- 返回err 返回error错误对象

功能说明：

这个函数主要是往一个文件里写入内容

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        b := make([]byte, 10)
        fi, err := os.OpenFile("Hello.go", os.O_RDWR | os.O_APPEND, 0420)
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        defer fi.Close()
        n, _ := fi.Read(b)
        fmt.Printf("len: %d, file content: %s\n", n, b[:n])
        fi.Write([]byte("cc\n"))
        fi.Seek(0, 0)
        n, _ = fi.Read(b)
        fmt.Printf("now len: %d, file content: %s\n", n, b[:n])
    }

代码输出：

    //test in ArchLinux
    len: 6, file content: aa
    bb

    now len: 9, file content: aa
    bb
    cc
