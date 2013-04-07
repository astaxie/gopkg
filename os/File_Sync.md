## func (f *File) Sync() (err error)

参数列表

- 无

返回值：

- 返回err 返回error错误对象

功能说明：

这个函数主要是把f内存里的内容写到磁盘上，即使是断电或系统崩溃，也能做到数据不丢失（注：一般不用调用这个方法）

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fi, err := os.OpenFile("Hello.go", os.O_RDWR | os.O_APPEND, 0420)
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        defer fi.Close()
        fi.Write([]byte("cc\n"))
        fi.Sync()
    }

代码输出：

    //test in ArchLinux
