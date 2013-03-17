## (c *Cmd) CombinedOutput() ([]byte, error)

参数列表

- 无

返回值：

- 返回[]byte 命令的执行结果
- 返回error 返回错误信息对象

功能说明：

这个函数主要是执行*Cmd中的命令，把执行结果和错误合并到byte数组中

代码实例：

    package main

    import (
        "fmt"
        "os/exec"
    )

    func main() {
        arg := []string{"Hello", "World!"}
        cmd := exec.Command("mv", arg...)
        output, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Printf("Error: %s\n", err)  //Error: exit status 1
        }
        //test in ArchLinux
        //当前目录中必须没有名为Hello的文件或文件夹
        fmt.Printf("The output is: %s\n", output)   //The output is: mv: 无法获取"Hello" 的文件状态(stat): 没有那个文件或目录
    }

