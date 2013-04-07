## func (e *LinkError) Error() string

参数列表

- 无

返回值：

- 返回string 返回LinkError的字符串形式

功能说明：

这个函数主要是获取LinkError的字符串形式

代码实例：

    package main

    import (
        "errors"
        "fmt"
        "os"
    )

    func main() {
        le := &os.LinkError{
            Op:  "ln",
            Old: "old",
            New: "new",
            Err: errors.New("ln Error!"),
        }
        fmt.Printf("%s\n", le.Error())
    }

代码输出：

    //test in ArchLinux
    ln old new: ln Error!
