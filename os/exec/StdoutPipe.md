## func (c *Cmd) StdoutPipe() (io.ReadCloser, error)

参数列表

- 无

返回值：

- 返回io.ReadCloser 连接到标准输出的管道
- 返回error 返回错误信息对象

功能说明：

这个函数主要是用于连接到命令启动时标准输出的管道，命令结束时，管道会自动关闭

代码实例：

    package main

    import (
        "fmt"
        "io/ioutil"
        "os/exec"
    )

    func main() {
        cmd := exec.Command("echo", "Hello World!")
        stdout, err := cmd.StdoutPipe()
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        if err = cmd.Start(); err != nil {
            fmt.Printf("Error: %s\n", err)
        }
        output, _ := ioutil.ReadAll(stdout)
        if err = cmd.Wait(); err != nil {
            fmt.Printf("Error: %s\n", err)
        }
        //test in ArchLinux
        fmt.Printf("The output is: %s\n", output)   //The output is: Hello World!
    }
