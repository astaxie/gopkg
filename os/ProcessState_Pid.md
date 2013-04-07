## func (p *ProcessState) Pid() int

参数列表

- 无

返回值：

- 返回int 已退出进程的pid

功能说明：

这个函数主要是获取已退出进程的pid

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
        fmt.Printf("Process pid is: %d\n", ps.Pid())
    }

代码输出：

    //test in ArchLinux
    /home/mirclesu/Programs/go
    Process pid is: 7767
