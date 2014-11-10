## func (f *FlagSet) Uint(name string, value uint, usage string) *uint

参数列表
- name string   flag名称
- value uint 变量默认值
- usage string 提示信息

返回值
- *uint 返回一个uint类型的flag值的地址

功能说明
- 为flag集合新增一个带默认值和提示语句的uint类型flag，返回对应值的地址

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
