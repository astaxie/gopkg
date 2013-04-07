## func (e *PathError) Error() string

参数列表

- 无

返回值：

- 返回string 返回PathError的字符串形式

功能说明：

这个函数主要是获取PathError的字符串形式

代码实例：

    package main

    import (
        "errors"
        "fmt"
        "os"
    )

    func main() {
        pe := &os.PathError{
            Op:   "cp",
            Path: "/oh/my/god",
            Err:  errors.New("path does'n exists!"),
        }
        fmt.Printf("%s\n", pe.Error())
    }

代码输出：

    //test in ArchLinux
    cp /oh/my/god: path does'n exists!
