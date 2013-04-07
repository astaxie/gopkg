## func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)

参数列表

- name 执行命令
- argv 命令参数
- attr 命令相关环境参数

返回值：

- 返回*Process 返回进程结构体指针
- 返回error 返回error错误对象信息

功能说明：

这个函数主要是启动一个进程

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        attr := &os.ProcAttr{
            Files: []*os.File{os.Stdin, os.Stdout},
            Env: os.Environ(),
        }
        p, err := os.StartProcess("/bin/pwd", []string{"/bin/pwd"}, attr)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        fmt.Printf("Process info: %+v\n", p)
    }

代码输出：

    //test in ArchLinux
    Process info: &{Pid:2596 handle:0 isdone:0}
