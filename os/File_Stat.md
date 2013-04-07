## func (f *File) Stat() (fi FileInfo, err error)

参数列表

- 无

返回值：

- 返回fi 返回文件的信息
- 返回err 返回error错误对象

功能说明：

这个函数主要是获取文件信息

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
        f, _ := fi.Stat()
        fmt.Printf("file info: %+v\n", f)
    }

代码输出：

    //test in ArchLinux
    file info: &{name:Hello.go size:6 mode:420 modTime:{sec:63500829488 nsec:441408659 loc:0x4dd468} sys:0xf84004c000}
