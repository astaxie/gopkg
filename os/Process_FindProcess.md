## func FindProcess(pid int) (p *Process, err error)

参数列表

- pid 进程pid

返回值：

- 返回p 返回进程结构体指针
- 返回err 返回error错误对象信息

功能说明：

这个函数主要是通过进程pid查找一个进程对象

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        p, err := os.FindProcess(12)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        fmt.Printf("Process info: %+v\n", p)
    }

代码输出：

    //test in ArchLinux
    Process info: &{Pid:12 handle:0 isdone:0}
