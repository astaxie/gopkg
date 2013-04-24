## func (m FileMode) IsDir() bool

参数列表

- 无

返回值：

- 返回bool

功能说明：

这个函数主要是判断一个文件模式是否为文件夹

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        f, err := os.Open("Hello.go")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        defer f.Close()
        fi, _ := f.Stat()
        fmt.Printf("%s is dir?: %t\n", fi.Name(), fi.Mode().IsDir())
    }

代码输出：

    //test in ArchLinux
    Hello.go is dir?: false
