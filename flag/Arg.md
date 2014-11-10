## func Arg(i int) string

参数列表
- i int 非flag命令行参数集合的索引

返回值
- string 第i＋1个非flag命令行参数的值

功能说明
- 获取第i＋1个非flag命令行参数的值，使用前需要先调用`flag.Parse()`解析flag,底层实际调用了全局变量CommandLine的Arg函数

代码示例1
    
    package main
    
    import (
        "flag"
        "fmt"
    )
    
    func main() {
        flag.Parse()
        fmt.Println(flag.Arg(0))
    }

执行结果
    
    //  ./test_arg abc
    abc
    
代码示例2
    
    package main
    
    import (
        "flag"
        "fmt"
    )
    
    var ip = flag.String("flag", "1234", "test for flags")
    
    func main() {
        flag.Parse()
        fmt.Println(flag.Arg(0))
    }

执行结果
    
    //  ./test_arg abc
    abc
    
    //  ./test_arg -flag abc def
    def
    