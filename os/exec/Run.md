## (c *Cmd) Run() error

参数列表

- 无

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是执行*Cmd中的命令，并且会等待命令执行完成，如果命令执行不成功，则返回错误信息

代码实例：

    package main

    import (
        "bytes"
        "fmt"
        "os/exec"
    )

    func main() {
        arg := []string{"Hello", "World!"}
        cmd := exec.Command("echo", arg...)
        var output bytes.Buffer
        cmd.Stdout = &output
        err := cmd.Run()
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        //test in ArchLinux
        fmt.Printf("The output is: %s\n", output.Bytes())   //The output is: Hello World!
    }

