## func SameFile(fi1, fi2 FileInfo) bool

参数列表

- fi1 文件信息
- fi2 文件信息

返回值：

- 返回bool 两个信息是否相同

功能说明：

这个函数主要是判断两个文件信息是否相同

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fi1, err := os.Stat("Hello.go")
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        fi2, err := os.Stat("World.go")
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        fmt.Printf("%t\n", os.SameFile(fi1, fi2))
        fmt.Printf("%t\n", os.SameFile(fi1, fi1))
    }

代码输出：

    //test in ArchLinux
    false
    true
