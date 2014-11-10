## func Uint(name string, value uint, usage string) *uint

参数列表
- name string   flag名称
- value uint 变量默认值
- usage string 提示信息

返回值
- *uint 返回一个uint类型的flag值的地址

功能说明
- 定义一个带默认值和提示语句的uint类型flag，返回对应值的地址

代码示例
        
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var uintFlag = flag.Uint("uint", 100, "help message for uint")
    
    func main() {
    	flag.Parse()
    	fmt.Println("uintFlag: ", *uintFlag)
    }


代码输出
            
//  ./testuint
uintFlag:  100

// ./testuint -uint 100
uintFlag:  100

// ./testuint -uint=100
uintFlag:  100

// ./testuint -uint    
flag needs an argument: -uint
Usage of ./testuint:
  -uint=100: help message for uint
