## func IsPathSeparator(c uint8) bool

参数列表

- c 判断字符

返回值：

- bool 是否为目录路径分割符

功能说明：

这个函数主要是判断一个字符是否为目录路径分割符

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("%t\n", os.IsPathSeparator('/'))
        fmt.Printf("%t\n", os.IsPathSeparator('.'))
    }

代码输出：

    //test in ArchLinux
    true
    false
