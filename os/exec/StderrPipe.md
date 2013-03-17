## func (c *Cmd) StderrPipe() (io.ReadCloser, error)

参数列表

- 无

返回值：

- 返回io.ReadCloser 连接到错误标准输出的管道
- 返回error 返回错误信息对象

功能说明：

这个函数主要是用于连接到命令启动时错误标准输出的管道，命令结束时，管道会自动关闭

代码实例：

    package main

    import (
        "fmt"
        "os/exec"
    )

    func main() {
        arg := []string{"Hello", "World!"}
        cmd := exec.Command("mv", arg...)
        stderr, err := cmd.StderrPipe()
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        if err = cmd.Start(); err != nil {
            fmt.Printf("Error: %s\n", err)
        }
        output, err := ioutil.ReadAll(stderr)
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        if err = cmd.Wait(); err != nil {
            fmt.Printf("Error: %s\n", err)  //Error: exit status 1
        }
        //test in ArchLinux
        //当前目录中必须没有名为Hello的文件或文件夹
        fmt.Printf("The mv command error is: %s\n", output)   //The mv command error is: mv: 无法获取"Hello" 的文件状态(stat): 没有那个文件或目录
    }

