## func (f *File) Readdir(n int) (fi []FileInfo, err error)

参数列表

- int 需要取出的文件信息数

返回值：

- 返回[]FileInfo 返回文件目录中文件的信息slice
- 返回error 返回error错误对象

功能说明：

这个函数主要是读取目录下的文件信息，如果参数n>0，则读取min(n, 文件夹下的文件数目)个文件信息，如果n<=0，则读取文件夹下所有的文件信息

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fi, err := os.Open("Hello")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        defer fi.Close()
        ff, _ := fi.Readdir(10)
        for i, f := range ff {
            fmt.Printf("fileinfo %d: %+v\n", i, f)
        }
    }

代码输出：

    //test in ArchLinux
    //文件夹下有Hello.go World.go两个文件
    fileinfo 0: &{name:World.go size:0 mode:420 modTime:{sec:63500830281 nsec:418121275 loc:0x4de468} sys:0xf84004e000}
    fileinfo 1: &{name:Hello.go size:0 mode:420 modTime:{sec:63500830275 nsec:878120955 loc:0x4de468} sys:0xf84004e090}
