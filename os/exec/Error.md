## func (e *Error) Error() string

参数列表

- 无

返回值：

- 返回string 返回执行失败命令的错误信息

功能说明：

这个函数主要是输出命令执行失败的错误信息

代码实例：

    package main

    import (
        "errors"
        "fmt"
        "os/exec"
    )

    func main() {
        e := exec.Error{
            Name: "mv",
            Err:  errors.New("无法获取\"Hello\" 的文件状态(stat): 没有那个文件或目录"),
        }
        //test in ArchLinux
        fmt.Printf("%s\n", e.Error()) //exec: "mv": 无法获取"Hello" 的文件状态(stat): 没有那个文件或目录
    }
