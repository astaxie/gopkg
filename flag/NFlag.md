## func NFlag() int

参数列表

返回值
- int 返回命令行参数里已解析到的flag个数

功能说明
- 获取命令行参数里解析成功的flag个数  

示例
        
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var stringFlag = flag.String("abc", "default", "help message")
    var intFlag = flag.Int("int", 123, "help message")
    
    var m float64
    
    func init() {
    	flag.Float64Var(&m, "float", 12.3, "help message")
    }
    
    func main() {
    	fmt.Println("before parse:", flag.NFlag())
    	flag.Parse()
    	fmt.Println("after parse", flag.NFlag())
    }

代码输出
        
    //  ./testnflag -abc 123 -int 0
    before parse: 0
    after parse 2
