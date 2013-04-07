## func (f *File) ReadAt(b []byte, off int64) (n int, err error)

参数列表

- []byte 存储读取内容的byte slice
- int64 开始读取内容的下标，单位是byte

返回值：

- 返回int 返回读取文件内容的字节长度
- 返回error 返回error错误对象

功能说明：

这个函数主要是从文件的某处开始读取内容到[]byte

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
        n, _ := fi.ReadAt(b, 3)
        fmt.Printf("%d\n", n)
        fmt.Printf("%s\n", b[:n])
    }

代码输出：

    //test in ArchLinux
    //Hello.go文件内容"aa\nbb\n"
    3
    bb
