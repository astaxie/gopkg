## func (f *FlagSet) Set(name, value string) error

参数列表
- name string flag名称
- value string flag值

功能说明
- 修改指定flag的值

返回值
- error 设置成功返回nil

功能说明
- 修改指定flag的值，成功返回nil

示例
        
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
    	testFlag  = myFlagSet.String("testFlag", "default value", defaultUsage)
    	uintVar   uint
    )
    
    func init() {
    	myFlagSet.UintVar(&uintVar, "uintVar", 30, defaultUsage)
    }

    func main() {
    	args := []string{
    		"-uintFlag", "24",
    		"-testFlag", "test",
    		"arg1",
    	}
    	myFlagSet.Parse(args)
    	if myFlagSet.Parsed() {
    		fmt.Println("before set:", *testFlag)
    		myFlagSet.Set("testFlag", "12345")
    		fmt.Println("after set", *testFlag)
    	}
    }
    
代码输出
        
    // ./testflagsetset                        
    before set: test
    after set 12345

