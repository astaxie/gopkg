## (c *Cmd) Start() error

参数列表

- 无

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是执行*Cmd中的命令，只是让命令开始执行，并不会等待命令执行完。

代码实例：

    package main

    import (
        "fmt"
        "os/exec"
    )

    func main() {
        cmd := exec.Command("sleep", "5")
        err := cmd.Start()
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("Waiting for command to finish...\n")
        err = cmd.Wait()
        fmt.Printf("Command finished with error: %v\n", err)
    }

