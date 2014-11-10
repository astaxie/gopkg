## func (f *FlagSet) PrintDefaults()

参数列表

返回值

功能说明
- 将所有已设置flag的默认值输出，默认为输出到标准出错，可通过SetOutput接口修改，默认值包括默认值，默认错误提示等。

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
