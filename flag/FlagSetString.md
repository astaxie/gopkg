## func (f *FlagSet) String(name string, value string, usage string) *string

参数列表
- name string   flag名称
- value string 变量默认值
- usage string 提示信息

返回值
- *string 返回一个string类型的flag值的地址

功能说明
- 为f增加一个带默认值和提示语句的string类型flag，返回对应值的地址

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
