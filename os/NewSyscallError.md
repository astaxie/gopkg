## func NewSyscallError(syscall string, err error) error

参数列表

- syscall 系统调用名
- err 错误信息对象

返回值：

- 返回error 返回SyscallError对象

功能说明：

这个函数主要是新建SyscallError对象

代码实例：

    package main

    import (
        "errors"
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("%s\n", os.NewSyscallError("custom", errors.New("something wrong")))
    }

代码输出：

    //test in ArchLinux
    custom: something wrong
