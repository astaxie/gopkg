## func (e *ExitError) Error() string

参数列表

- 无

返回值：

- 返回string 返回一个执行不成功命令的信息

功能说明：

这个函数主要是返回一个执行不成功命令的信息

代码实例：

    package main

    import (
        "fmt"
        "os/exec"
    )

    func main() {
        cmd := exec.Command("mv", "Hello World!")
        cmd.Run()
        exitError := exec.ExitError{cmd.ProcessState}
        //test in ArchLinux
        fmt.Printf("The output is: %s\n", exitError.Error())   //The output is: exit status 1
    }
