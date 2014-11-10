## func Uint64Var(p *uint64, name string, value uint64, usage string)

参数列表
- p *uint64 需要与flag参数值绑定的变量地址
- name string  flag名称
- value uint64 变量默认值
- usage string 提示信息

返回值

功能说明
- 将命令行指定flag参数值绑定到uint64变量

示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var uint64varFlag uint64
    
    func init() {
    	flag.Uint64Var(&uint64varFlag, "flag", 123, "help message for uint64var")
    }
    
    func main() {
    	flag.Parse()
    	fmt.Println("uint64varFlag:", uint64varFlag)
    }


代码输出
        
    //./testuint64var 
    uint64varFlag: 123
    
    //  ./testuint64var -flag 
    flag needs an argument: -flag
    Usage of ./testuint64var:
      -flag=123: help message for uint64var
    
    //  ./testuint64var -flag 1234
    uint64varFlag: 1234
    
    //  ./testuint64var -flag=11234  
    uint64varFlag: 11234
