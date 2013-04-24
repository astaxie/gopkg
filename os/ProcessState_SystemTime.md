## func (p *ProcessState) SystemTime() time.Duration

参数列表

- 无

返回值：

- 返回time.Duration 进程系统cpu使用时间

功能说明：

这个函数主要是获取的进程的系统cpu使用时间

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
        fmt.Printf("Process system cpu time: %+v\n", ps.SystemTime())
    }

代码输出：

    //test in ArchLinux
    /home/mirclesu/Programs/go
    Process system cpu time : 0
