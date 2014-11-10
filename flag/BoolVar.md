## func BoolVar(p *bool, name string, value bool, usage string)

参数列表
- p *bool bool型变量指针
- name string flag名称
- value bool 默认值
- usage string 提示信息

返回值

功能说明
－ 将指定flag值绑定到一个bool变量，底层实际调用了全局变量CommandLine的BoolVar函数

代码示例
        
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var boolFlag bool
    
    func init() {
    	flag.BoolVar(&boolFlag, "flag", false, "help message for flag")
    }
    
    func main() {
    	flag.Parse()
    	fmt.Println("boolFlag: ", boolFlag)
    }

执行结果
    
    //  ./test_boolval
    boolFlag:  false
    
    //  ./test_boolval -flag
    boolFlag:  true
    
    //  ./test_boolval -flag=true
    boolFlag:  true