## func (f *FlagSet) StringVar(p *string, name string, value string, usage string)

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
    
    const (
    	defaultUsage = "help message"
    )
    
    var (
    	myFlagSet  = flag.NewFlagSet("newflagset", flag.ExitOnError)
    	stringFlag = myFlagSet.String("stringFlag", "defaultValue", defaultUsage)
    	stringVar  string
    )
    
    func init() {
    	myFlagSet.StringVar(&stringVar, "stringVar", "defaultValue", defaultUsage)
    }
    
    func main() {
    	args := []string{
    		"--stringFlag", "test string flag",
    		"-stringVar", "test string",
    		"arg2",
    	}
    	myFlagSet.Parse(args)
    	fmt.Println("stringFlag", *stringFlag)
    	fmt.Println("stringVar", stringVar)
    }

代码输出
        
    //  ./testflagsetstring 
    stringFlag test string flag
    stringVar test string
