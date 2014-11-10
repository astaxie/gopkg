## func Parse()

参数列表

返回值

功能说明
- 从os.Args[1:]中依次解析flag.在定义flag之后，必须调用用本函数才能成功获取命令行flag的值 

示例
        
    package main
    
    import (
        "flag"
        "fmt"
    )
    
    var stringFlag = flag.String("test", "test", "help message.")
    var intFlag = flag.Int("int", 10, "help message")
    
    func main() {
        fmt.Println("before parse:", *stringFlag)
        fmt.Println("before parse:", *intFlag)
        flag.Parse()
        fmt.Println("after parse:", *stringFlag)
        fmt.Println("after parse:", *intFlag)
    }

代码输出
        
    //  ./testparse -test "good" -int 100
    before parse: test
    before parse: 10
    after parse: good
    after parse: 100

