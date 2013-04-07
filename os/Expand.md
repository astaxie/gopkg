## func Expand(s string, mapping func(string) string) string

参数列表

- s 需要替换的字符串
- mapping 替换规则函数

返回值：

- string 替换后的字符串

功能说明：

这个函数主要是把一个字符串里带$var或${var}这样的字符串按自定义的规则替换成自己想要的字符串，如果没有对应的，则替换成空

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        mapping := func(s string) string {
            m := map[string]string{"hello":"world", "go":"perfect program language"}
            return m[s]
        }
        str := "Golang is$not a $go in the ${hello}!"
        fmt.Printf("%s\n", os.Expand(str, mapping))
    }

代码输出：

    //test in ArchLinux
    Golang is a perfect program language in the world!
