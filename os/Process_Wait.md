## func (p *Process) Wait() (*ProcessState, error)

参数列表

- 无

返回值：

- 返回*ProcessState 返回进程状态结构体指针
- 返回error 返回error错误对象信息

功能说明：

这个函数主要是等待一个进程执行完成

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
        fmt.Printf("Process stat: %+v\n", ps)
    }

代码输出：

    //test in ArchLinux
    /home/mirclesu/Programs/go
    Process stat: exit status 0
