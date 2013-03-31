## func (c *Cmd) StdinPipe() (io.WriteCloser, error)

参数列表

- 无

返回值：

- 返回io.WriteCloser 连接到标准输入的管道
- 返回error 返回错误信息对象

功能说明：

这个函数主要是用于连接到命令启动时标准输入的管道

代码实例：

    package main

    import (
        "bytes"
        "fmt"
        "os/exec"
    )

    func main() {
        var output bytes.Buffer
        cmd := exec.Command("cat")
        cmd.Stdout = &output
        stdin, err := cmd.StdinPipe()
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        if err = cmd.Start(); err != nil {
            fmt.Printf("Error: %s\n", err)
        }
        stdin.Write([]byte("Hello World!"))
        stdin.Close()
        if err = cmd.Wait(); err != nil {
            fmt.Printf("Error: %s\n", err)
        }
        //test in ArchLinux
        fmt.Printf("The output is: %s\n", output.Bytes())   //The output is: Hello World!
    }

