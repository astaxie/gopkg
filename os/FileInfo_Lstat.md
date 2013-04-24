## func Lstat(name string) (fi FileInfo, err error)

参数列表

- name 文件名

返回值：

- 返回fi 文件信息结构体
- 返回err 返回error错误对象

功能说明：

这个函数主要是获取一个文件的信息，如果这个文件是一个链接，返回的是链接本身的信息

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fi, err := os.Lstat("Hello.go")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("file info: %+v\n", fi)
    }

代码输出：

    //test in ArchLinux
    file info: &{name:Hello.go size:6 mode:420 modTime:{sec:63500837998 nsec:888568469 loc:0x4dd468} sys:0xf84004c000}
