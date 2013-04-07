## func IsNotExist(err error) bool

参数列表

- err 错误信息对象

返回值：

- bool 是否为文件或目录不存在的错误

功能说明：

这个函数主要是判断一个错误是否为文件或目录不存在的错误

代码实例：

    package main

    import (
        "errors"
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("%t\n", os.IsNotExist(os.ErrNotExist))
        fmt.Printf("%t\n", os.IsNotExist(errors.New("Custom Error")))
    }

代码输出：

    //test in ArchLinux
    true
    false
