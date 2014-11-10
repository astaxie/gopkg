## func (f *FlagSet) Parsed() bool

参数列表

返回值
- bool 若flag设置参数已解析(已调用f.Parse()方法)，返回true

功能说明
- 判断flag设置参数是否已经解析(是否已调用f.Parse()方法)，若已解析返回true

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
    		fmt.Println("uintFlag:", *uintFlag)
    		fmt.Println("testFlag:", *testFlag)
    	}
    }

代码输出
        
    //  ./testflagsetparsed                          
    uintFlag: 24
    testFlag: test
