## func Int64(name string, value int64, usage string) *int64

参数列表
- name string   flag名称
- value int64 变量默认值
- usage string 提示信息

返回值
- *int64 返回一个int64类型的flag值的地址

功能说明
- 定义一个带默认值和提示语句的int64类型flag，返回对应值的地址

代码示例
        
    package main
    
    import (
        "flag"
        "fmt"
    )
    
    var int64Flag = flag.Int64("int64", int64(100), "help message for int64")
    
    func main() {
        flag.Parse()
        fmt.Println("intFlag: ", *int64Flag)
    }

代码输出
            
     ./testint64
    intFlag:  100
    
    //  ./testint64 -int64 111111111111
    intFlag:  111111111111
    
    //  ./testint64 -int64=-123123121231
    intFlag:  -123123121231
    
    //  ./testint64 -int64              
    flag needs an argument: -int64
    Usage of ./testint64:
      -int64=100: help message for int64
