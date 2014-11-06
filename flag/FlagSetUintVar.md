## func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)

参数列表
- p *uint 需要与flag参数值绑定的变量地址
- name string  flag名称
- value uint 变量默认值
- usage string 提示信息

返回值

功能说明
- 将f中指定flag参数值绑定到uint变量

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
    	myFlagSet = flag.NewFlagSet("newflagset", flag.ExitOnError)
    	uintFlag  = myFlagSet.Uint("uintFlag", 1000, defaultUsage)
    	uintVar   uint
    )
    
    func init() {
    	myFlagSet.UintVar(&uintVar, "uintVar", 30, defaultUsage)
    }
    
    func main() {
    	args := []string{
    		"-uintFlag", "24",
    		"--uintVar", "25",
    		"arg2",
    	}
    	myFlagSet.Parse(args)
    	fmt.Println("uintFlag", *uintFlag)
    	fmt.Println("uintVar", uintVar)
    }

代码输出
            
    //  ./testflagsetuint                         
    uintFlag 24
    uintVar 25
