## func (p *ProcessState) Exited() bool

参数列表

- 无

返回值：

- 返回bool 进程是否已退出

功能说明：

这个函数主要是判断一个进程是否已退出

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
        fmt.Printf("Process has been exited?: %t\n", ps.Exited())
    }

代码输出：

    //test in ArchLinux
    /home/mirclesu/Programs/go
    Process has been exited?: true
