## func Duration(name string, value time.Duration, usage string) *time.Duration

参数列表
- name string flag名称
- value time.Duration 默认值
- usage string 提示信息 

返回值
- *time.Duration 返回一个Duration类型的flag值的地址

功能说明
- 定义一个带默认值和提示语句的Duration类型flag，返回flag对应值的地址

代码示例
    
    package main
    
    import (
        "flag"
        "fmt"
    )
    
    var durationFlag = flag.Duration("duraT", 12, "help message for duraT")
    
    func main() {
        flag.Parse()
        fmt.Println(*durationFlag)
    }

代码输出
    
    //  ./testduration
    12ns
    
    //  ./testduration -duraT
    flag needs an argument: -duraT
    Usage of ./testduration:
      -duraT=12ns: help message for duraT
      
    //  ./testduration -duraT=200ms
    200ms
    
    //  ./testduration -duraT 3000h
    3000h0m0s
