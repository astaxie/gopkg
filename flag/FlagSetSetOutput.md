## func (f *FlagSet) SetOutput(output io.Writer)

参数列表
- output io.Writer 设置错误信息和帮助信息的输出，默认为输出到标准错误

返回值

功能说明
- 设置错误信息和帮助信息的输出方式，默认输出到标准错误

示例
        
    package main
    
    import (
    	"flag"
    	"fmt"
    	"os"
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
    		myFlagSet.SetOutput(os.Stderr)
    		fmt.Println("PrintDefaults:")
    		myFlagSet.PrintDefaults()
    	}
    }

代码输出
        
    //  ./testflagsetprintdefaults
    PrintDefaults:
      -testFlag="default value": help message
      -uintFlag=1000: help message
      -uintVar=30: help message
