## func Command(name string, arg ...string) *Cmd

参数列表

- name 表示需要执行的命令名
- arg  表示跟在命令后的参数列表

返回值：

- 返回*Cmd 按参数初始化Path和Args字段，其他字段执行默认初始化的Cmd结构体指针

功能说明：

这个函数主要是用来初始化一个Cmd指针，Path和Args按参数初始化，其他字段执行默认初始化

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

