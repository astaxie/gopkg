## func NArg() int

参数列表

返回值
- int 返回解析后剩余的参数的数量

功能说明
- 获取命令行参数解析后剩余的参数的数量，即flag.Args()返回的值的元素个数   

示例
        
    package main
    
    import (
        "flag"
        "fmt"
    )
    
    var test_flag = flag.String("flag", "test", "help message for flag")
    
    func main() {
        flag.Parse()
        fmt.Println("test_flag:", *test_flag)
        fmt.Println("Args:", flag.Args())
        fmt.Println("Narg", flag.NArg())
    }

代码输出
        
    ./testnarg
    test_flag: test
    Args: []
    Narg 0
    
    //  ./testnarg asd
    test_flag: test
    Args: [asd]
    Narg 1
    
    //  ./testnarg -flag 123
    test_flag: 123
    Args: []
    Narg 0
    
    //  ./testnarg -flag 123 456 678 111
    test_flag: 123
    Args: [456 678 111]
    Narg 3
