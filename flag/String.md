## func String(name string, value string, usage string) *string

参数列表
- name string   flag名称
- value string 变量默认值
- usage string 提示信息

返回值
- *string 返回一个string类型的flag值的地址

功能说明
- 定义一个带默认值和提示语句的string类型flag，返回对应值的地址

代码示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var stringFlag = flag.String("string", "default value", "help message for int")
    
    func main() {
    	flag.Parse()
    	fmt.Println("stringFlag: ", *stringFlag)
    }


代码输出
        
    //  ./teststring 
    stringFlag:  default value
    
    //  ./teststring -string
    flag needs an argument: -string
    Usage of ./teststring:
      -string="default value": help message for int
    
    //  ./teststring -string=123
    stringFlag:  123
    
    //  ./teststring -string 123
    stringFlag:  123

