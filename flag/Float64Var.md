## func Float64Var(p *float64, name string, value float64, usage string)

参数列表
- p *float64 需要与flag参数值绑定的变量地址
- name string   flag名称
- value float64 变量默认值
- usage string 提示信息

返回值

功能说明
- 将命令行指定flag参数值绑定到float64变量

代码示例
    
    package main
    
    import (
        "flag"
        "fmt"
    )
    
    var float64Var float64
    
    func init() {
        flag.Float64Var(&float64Var, "float64", 123.4, "help message for float64")
    }
    
    func main() {
        flag.Parse()
        fmt.Println("float64Flag:", float64Var)
    }


代码输出
    
    //  ./testfloat64var 
    float64Flag: 123.4
    
    //  ./testfloat64var -float64 12345
    float64Flag: 12345
    
    //  ./testfloat64var -float64=1231.1
    float64Flag: 1231.1
    
    //  ./testfloat64 -float64
    flag needs an argument: -float64
    Usage of ./testfloat64:
      -float64=123.4: help message for float64