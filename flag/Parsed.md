## func Parsed() bool

参数列表

返回值
- 若命令行flag参数已解析(已调用flag.Parse()方法)，返回true

功能说明
- 判断命令行flag参数是否已经解析(即调用flag.Parse()方法)，若已解析返回true

示例
        
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var stringFlag = flag.String("test", "test", "help message.")
    var intFlag = flag.Int("int", 10, "help message")
    
    func main() {
    	if flag.Parsed() {
    		fmt.Println("before parse:", *stringFlag)
    		fmt.Println("before parse:", *intFlag)
    
    	}
    	flag.Parse()
    	if flag.Parsed() {
    		fmt.Println("after parse:", *stringFlag)
    		fmt.Println("after parse:", *intFlag)
    	}
    }

代码输出
        
    ./testparsed -test abc
    after parse: abc
    after parse: 10


