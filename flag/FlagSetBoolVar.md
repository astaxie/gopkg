## func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)

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

