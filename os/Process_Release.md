## func (p *Process) Release() error

参数列表

- 无

返回值：

- 返回error 返回error错误对象信息

功能说明：

这个函数主要是释放一个进程相关的资源

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
        if err := p.Release(); err != nil {
            fmt.Printf("Error: %v\n", err)
        }
    }

代码输出：

    //test in ArchLinux
    Process info: &{Pid:3997 handle:0 isdone:0}
