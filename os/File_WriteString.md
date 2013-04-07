## func (f *File) WriteString(s string) (ret int, err error)

参数列表

- s 要写入的内容

返回值：

- 返回ret 返回写入的字节数
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
        b := make([]byte, 20)
        fi, err := os.OpenFile("Hello.go", os.O_RDWR | os.O_APPEND, 0420)
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        defer fi.Close()
        n, _ := fi.Read(b)
        fmt.Printf("len: %d, file content: %s\n", n, b[:n])
        fi.WriteString("Hello World!\n")
        fi.Seek(0, 0)
        n, _ = fi.Read(b)
        fmt.Printf("now len: %d, file content: %s\n", n, b[:n])
    }

代码输出：

    //test in ArchLinux
    len: 6, file content: aa
    cc

    now len: 9, file content: aa
    bb
    Hello World!
