## func Bool(name string, value bool, usage string) *bool

参数列表
- name string 命令行flag名称
- value bool 默认值
- usage string 帮助信息

返回值
- *bool 返回一个bool类型的flag值的地址

功能说明
- 定义一个带默认值和提示语句的bool类型flag，返回flag对应值的地址，底层实际调用了全局变量CommandLine的Bool方法

代码示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var boolFlag = flag.Bool("flag", false, "help message for flag")
    
    func main() {
    	flag.Parse()
    	fmt.Println("boolFlag: ", *boolFlag)
    }

执行结果
    
    //  ./test_bool
    boolFlag:  false
    
    //  ./test_bool -flag
    boolFlag:  true
    
    //  ./test_bool -flag=false
    boolFlag:  false
