## func LookPath(file string) (string, error)

参数列表

- file 表示环境变量中执行的二进制文件名

返回值：

- 返回string 可执行文件所在的路径
- 返回error 返回错误信息对象

功能说明：

这个函数主要是用来查询可执行二进制文件的路径

代码实例：

    package main

    import (
        "fmt"
        "os/exec"
    )

    func main() {
        path, err := exec.LookPath("echo")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        //测试环境为ArchLinux
        fmt.Printf("echo is available at %s\n", path)    //echo is available at /usr/bin/echo
    }

