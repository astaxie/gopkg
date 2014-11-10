## func Args() []string

参数列表

返回值
- []string 非flag命令行参数集合

功能说明
- 获取非flag命令行参数集合，使用前需要先调用`flag.Parse`解析flag，底层实际调用了全局变量CommandLine的Args函数

代码示例1
    
    package main
    
    import (
        "flag"
        "fmt"
    )
    
    func main() {
        flag.Parse()
        fmt.Println(flag.Args())
    }


执行结果
    
    //  ./test_args abc
    [abc]
    
代码示例2
    
    package main
    
    import (
        "flag"
        "fmt"
    )
    
    var ip = flag.String("flag", "1234", "test for flags")
    
    func main() {
        flag.Parse()
        fmt.Println(flag.Args())
    }

执行结果
    
    //  ./test_args abc def
    [abc def]
    
    //  ./test_args -flag abc def ghi
    [def ghi]
    