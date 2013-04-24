## func Chtimes(name string, atime time.Time, mtime time.Time) error

参数列表

- name  修改的文件名
- atime 访问时间
- mtime 修改时间

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是修改一个文件的访问时间和修改时间

代码实例：

    package main

    import (
        "fmt"
        "os"
        "syscall"
        "time"
    )

    func main() {
        fi, err := os.Stat("Hello.go")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("Hello.go: atime=%+v, mtime=%v\n", fi.Sys().(*syscall.Stat_t).Atim, fi.ModTime())

        err = os.Chtimes("Hello.go", time.Now(), time.Now())
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }

        fi, err = os.Stat("Hello.go")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("Now Hello.go: atime=%+v, mtime=%v\n", fi.Sys().(*syscall.Stat_t).Atim, fi.ModTime())
    }

代码输出：

    //test in ArchLinux
    Hello.go: atime={Sec:1363663572 Nsec:515975000}, mtime=2013-03-19 11:26:12.515978 +0800 CST
    Now Hello.go: atime={Sec:1363663747 Nsec:130183000}, mtime=2013-03-19 11:29:07.130184 +0800 CST
