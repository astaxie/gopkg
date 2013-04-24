## func (p *ProcessState) Sys() interface{}

参数列表

- 无

返回值：

- 返回interface{}

功能说明：

这个函数主要是获取进程的退出信息

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
        ps, _ := p.Wait()
        fmt.Printf("Process exit status: %+v\n", ps.Sys())
    }

代码输出：

    //test in ArchLinux
    /home/mirclesu/Programs/go
    Process exit status: 0
