## func UintVar(p *uint, name string, value uint, usage string)

参数列表
- p *uint 需要与flag参数值绑定的变量地址
- name string  flag名称
- value uint 变量默认值
- usage string 提示信息

返回值

功能说明
- 将命令行指定flag参数值绑定到uint变量

代码示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var uintvarFlag uint
    
    func init() {
    	flag.UintVar(&uintvarFlag, "flag", 123, "help message for uintvar")
    }
    
    func main() {
    	flag.Parse()
    	fmt.Println("uintvarFlag:", uintvarFlag)
    }

代码输出
        
    //./testuintvar 
    uintvarFlag: 123
    
    // ./testuintvar -flag  
    flag needs an argument: -flag
    Usage of ./testuintvar:
      -flag=123: help message for uintvar
    
    // ./testuintvar -flag 456
    uintvarFlag: 456
    
    // ./testuintvar -flag=456
    uintvarFlag: 456
