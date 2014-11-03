## func Uint64(name string, value uint64, usage string) *uint64

参数列表
- name string   flag名称
- value uint64 变量默认值
- usage string 提示信息

返回值
- *uint64 返回一个uint64类型的flag值的地址

功能说明
- 定义一个带默认值和提示语句的uint64类型flag，返回对应值的地址

代码示例
        
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var uint64Flag = flag.Uint64("uint64", 100, "help message for uint")
    
    func main() {
    	flag.Parse()
    	fmt.Println("uint64Flag: ", *uint64Flag)
    }

代码输出
            
//  ./testuint64                    
uint64Flag:  100

//  ./testuint64 -uint64            
flag needs an argument: -uint64
Usage of ./testuint64:
  -uint64=100: help message for uint

//  ./testuint64 -uint64 1231312    
uint64Flag:  1231312

//  ./testuint64 -uint64=1231312    
uint64Flag:  1231312
