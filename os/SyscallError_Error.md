## func (e *SyscallError) Error() string

参数列表

- 无

返回值：

- 返回string 返回SyscallError的字符串形式

功能说明：

这个函数主要是获取SyscallError的字符串形式

代码实例：

    package main

    import (
        "errors"
        "fmt"
        "os"
    )

    func main() {
        se := &os.SyscallError{
            Syscall: "/xx",
            Err:     errors.New("Syscall Error!"),
        }
        fmt.Printf("%s\n", se.Error())
    }

代码输出：

    //test in ArchLinux
    /xx: Syscall Error!
