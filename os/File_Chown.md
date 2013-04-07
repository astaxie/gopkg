## func (f *File) Chown(uid, gid int) error

参数列表

- uid  用户id
- gid  用户组id

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是修改文件的所属用户和用户组

代码实例：

    package main

    import (
        "fmt"
        "os"
        "syscall"
    )

    func main() {
        fi, err := os.Stat("Hello.go")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("Hello.go: uid=%d, gid=%d\n", fi.Sys().(*syscall.Stat_t).Uid, fi.Sys().(*syscall.Stat_t).Gid)

        f, _ := os.Open("Hello.go")
        err = f.Chown(99, 99)  //nobody, nobody
        if err != nil {
            fmt.Printf("Error: %s\n", err)
        }
        f.Close()

        fi, err = os.Stat("Hello.go")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("Now Hello.go: uid=%d, gid=%d\n", fi.Sys().(*syscall.Stat_t).Uid, fi.Sys().(*syscall.Stat_t).Gid)
    }

代码输出：

    //test in ArchLinux by root user
    Hello.go: uid=0, gid=0
    Now Hello.go: uid=99, gid=99
