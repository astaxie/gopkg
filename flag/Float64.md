## func Float64(name string, value float64, usage string) *float64

参数列表

- name string flag名称
- value float64 默认值
- usage string 提示信息

返回值
- *float64 返回一个float64类型的flag值的地址

功能说明
- 定义一个带默认值和语句的float64类型的flag，返回flag对应值的地址

代码示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var float64Flag = flag.Float64("float64", 123.4, "help message for float64")
    
    func main() {
    	flag.Parse()
    	fmt.Println("float64Flag: ", *float64Flag)
    }

代码输出
        
    //  ./testfloat64 
    float64Flag:  123.4

    //  ./testfloat64 -float64 1230.6
    float64Flag:  1230.6
    
    //  ./testfloat64 -float64=0.123
    float64Flag:  0.123

