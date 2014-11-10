## func Int64Var(p *int64, name string, value int64, usage string)

参数列表
- p *int64 需要与flag参数值绑定的变量地址
- name string  flag名称
- value int64 变量默认值
- usage string 提示信息

返回值

功能说明
- 将指定flag参数值绑定到int64变量    

示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var int64Flag int64
    
    func init() {
    	flag.Int64Var(&int64Flag, "int64", 123, "help message for int64")
    }
    
    func main() {
    	flag.Parse()
    	fmt.Println("int64Flag:", int64Flag)
    }

代码输出
    
    //  ./testint64var 
    int64Flag: 123
    
    //  ./testint64var -int64
    flag needs an argument: -int64
    Usage of ./testint64var:
      -int64=123: help message for int64
    
    //  ./testint64var -int64=12312312
    int64Flag: 12312312
    
    //  ./testint64var -int64 12312312
    int64Flag: 12312312
