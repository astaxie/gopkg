## func (p *Process) Signal(sig Signal) error

参数列表

- sig 系统信号

返回值：

- 返回error 返回error错误对象信息

功能说明：

这个函数主要是给一个进程发送一个信号

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
        if err := p.Signal(os.Kill); err != nil {
            fmt.Printf("Error: %v\n", err)
        }
    }

代码输出：

    //test in ArchLinux
    Process info: &{Pid:4497 handle:0 isdone:0}
