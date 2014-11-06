## func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration

参数列表
- name string flag名称
- value time.Duration 默认值
- usage string 提示信息 

返回值
- *time.Duration 返回一个Duration类型的flag值的地址

功能说明
- 为Flag集合f增加一个带默认值和提示语句的Duration类型flag，返回flag对应值的地址

代码示例
        
    package main
    
    import (
    	"flag"
    	"fmt"
    	"time"
    )
    
    const (
    	defaultUsage = "help message"
    )
    
    var (
    	myFlagSet = flag.NewFlagSet("newflagset", flag.ExitOnError)
    	duratFlag = myFlagSet.Duration("duratFlag", 20, defaultUsage)
    	duratVar  time.Duration
    )
    
    func init() {
    	myFlagSet.DurationVar(&duratVar, "duratVar", 100, defaultUsage)
    }
    
    func main() {
    	args := []string{
    		"-duratVar", "2m",
    		"-duratFlag", "12h",
    		"arg2",
    	}
    	myFlagSet.Parse(args)
    	fmt.Println("duratFlag", *duratFlag)
    	fmt.Println("duratVar", duratVar)
    }

代码输出
    
    //  ./testflagsetduration 
    duratFlag 12h0m0s
    duratVar 2m0s

