## func Set(name, value string) error

参数列表
- name string flag 名称
- value string flag 值

返回值
- error 设置成功返回nil

功能说明
- 将名称为name的flag的值设置为value, 成功返回nil

示例
        
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var stringFlag = flag.String("test", "test", "help message.")
    var intFlag = flag.Int("int", 10, "help message")
    
    func main() {
    	flag.Parse()
    	fmt.Println("stringFlag:", *stringFlag)
    	fmt.Println("inFlag", *intFlag)
    	fmt.Println("After Set.")
    	flag.Set("test", "set_test")
    	flag.Set("int", "9999")
    	fmt.Println("stringFlag:", *stringFlag)
    	fmt.Println("inFlag:", *intFlag)
    }

代码输出
        
    //  ./testset -test "before set" -int 1
    stringFlag: before set
    inFlag: 1
    After Set.
    stringFlag: set_test
    inFlag: 9999
