## func Size(v interface{}) int

###参数列表

- v 需要计算长度的数据

###返回值：

- int 数据序列化之后的字节长度

###功能说明：

Size将返回数据序列化之后的字节长度，数据必须是固定长数据类型、slice和结构体及其指针等。

###代码实例：
    package main
    
    import (
      "encoding/binary"
    	"fmt"
    )
    
    func main() {
    	var a int
    	b := [5]int64{1}
    
    	fmt.Println(binary.Size(a)) //-1
    	fmt.Println(binary.Size(b)) //40
    }
