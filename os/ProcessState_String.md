## func (p *ProcessState) String() string

参数列表

- 无

返回值：

- 返回string 进程状态的字符串

功能说明：

这个函数主要是获取进程状态的字符串

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
        fmt.Printf("Process stat: %s\n", ps.String())
    }

代码输出：

    //test in ArchLinux
    /home/mirclesu/Programs/go
    Process stat: exit status 0
