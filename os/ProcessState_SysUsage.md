## func (p *ProcessState) SysUsage() interface{}

参数列表

- 无

返回值：

- 返回interface{}

功能说明：

这个函数主要是获取进程资源使用情况

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
        fmt.Printf("Process resource usage: %+v\n", ps.SysUsage())
    }

代码输出：

    //test in ArchLinux
    /home/mirclesu/Programs/go
    Process resource usage: &{Utime:{Sec:0 Usec:0} Stime:{Sec:0 Usec:0} Maxrss:640 Ixrss:0 Idrss:0 Isrss:0 Minflt:179 Majflt:0 Nswap:0 Inblock:0 Oublock:0 Msgsnd:0 Msgrcv:0 Nsignals:0 Nvcsw:1 Nivcsw:3}
