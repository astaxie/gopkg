## func DurationVar(p *time.Duration, name string, value time.Duration, usage string)

参数列表
- p *time.Duration  需要与flag值绑定的变量的指针
- name string   flag名称
- value time.Duration   默认值
- usage string  提示信息

返回值

功能说明
- 将一个time.Duration类型的flag值绑定到一个time.Duration的变量 

代码示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    	"time"
    )
    
    var durationFlag time.Duration
    
    func init() {
    	flag.DurationVar(&durationFlag, "duraT", 120, "help message for duraT")
    }
    
    func main() {
    	flag.Parse()
    	fmt.Println(durationFlag)
    }

代码输出
    
    //  ./testdurationvar
    120ns
    
    //  ./testdurationvar -duraT
    flag needs an argument: -duraT
    Usage of ./testdurationvar:
      -duraT=120ns: help message for duraT
      
    //  ./testdurationvar -duraT=200ms
    200ms
    
    //  ./testdurationvar -duraT 3000h
    3000h0m0s