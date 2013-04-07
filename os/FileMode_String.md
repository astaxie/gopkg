## func (m FileMode) String() string

参数列表

- 无

返回值：

- 返回string 返回文件模式的字符串形式

功能说明：

这个函数主要是获取文件模式的字符串形式

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
        fmt.Printf("%s mode: %v\n", fi.Name(), fi.Mode().String())
    }

代码输出：

    //test in ArchLinux
    Hello.go mode: -rw-r--r--
