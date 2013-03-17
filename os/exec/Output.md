## (c *Cmd) Output() ([]byte, error)

参数列表

- 无

返回值：

- 返回[]byte 命令的执行结果
- 返回error 返回错误信息对象

功能说明：

这个函数主要是执行*Cmd中的命令，返回执行结果

代码实例：

    package main

    import (
        "fmt"
        "os/exec"
    )

    func main() {
        arg := []string{"Hello", "World!"}
        cmd := exec.Command("echo", arg...)
        output, err := cmd.Output()
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        //test in ArchLinux
        fmt.Printf("The output is: %s\n", output)   //The output is: Hello World!
    }

