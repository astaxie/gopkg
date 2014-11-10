## func IntVar(p *int, name string, value int, usage string)

参数列表
- p *int 需要与flag参数值绑定的变量地址
- name string  flag名称
- value int 变量默认值
- usage string 提示信息

返回值

功能说明
- 将命令行指定flag参数值绑定到int变量

示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var intFlag int
    
    func init() {
    	flag.IntVar(&intFlag, "int", 123, "help message for int")
    }
    
    func main() {
    	flag.Parse()
    	fmt.Println("intFlag:", intFlag)
    }

代码输出
    
    //  ./testintvar 
    intFlag: 123
    
    // ./testintvar -int 456
    intFlag: 456
    
    // ./testintvar -int=456
    intFlag: 456
    
    // ./testintvar -int    
    flag needs an argument: -int
    Usage of ./testintvar:
      -int=123: help message for int
