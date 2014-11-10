## func PrintDefaults()

参数列表

返回值

功能说明
- 将所有已定义flag的默认值输出到标准错误，默认值包括默认值，默认错误提示等。

示例
        
    package main
    
    import (
    	"flag"
    )
    
    var stringFlag = flag.String("test", "test", "help message.")
    var intFlag = flag.Int("int", 10, "help message")
    
    func main() {
    	flag.PrintDefaults()
    }

代码输出
        
    //  ./testprintdefaults        
    -int=10: help message
    -test="test": help message.
