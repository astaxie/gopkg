## func StringVar(p *string, name string, value string, usage string)

参数列表
- p *string 需要与flag参数值绑定的变量地址
- name string   flag名称
- value string 变量默认值
- usage string 提示信息

返回值

功能说明
- 将命令行指定flag参数值绑定到string变量

代码示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var stringFlag string
    
    func init() {
    	flag.StringVar(&stringFlag, "flag", "default value", "help message for flag")
    }
    
    func main() {
    	flag.Parse()
    	fmt.Println("flag: ", stringFlag)
    }
 
代码输出
    
    //  ./teststringvar 
    flag:  default value
    
    //  ./teststringvar -flag    
    flag needs an argument: -flag
    Usage of ./teststringvar:
      -flag="default value": help message for flag
    
    // ./teststringvar -flag=12
    flag:  12
    
    // ./teststringvar -flag 22
    flag:  22
