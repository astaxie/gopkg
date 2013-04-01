## Write(w io.Writer, order ByteOrder, data interface{}) error

###参数列表

- w 可写入字节流的数据 
- order 特殊字节序，包中提供大端字节序和小段字节序
- data 需要解码成的数据

###返回值：

- error 返回错误

###功能说明：

Write将data序列化成字节流写入w中。data必须是固定长的数据值或固定长数据的slice，或指向此类数据的指针。写入w的字节流可用特殊的字节序来编码。另外，结构体中的(_)名的字段将忽视。

###代码实例：
    package main
    
    import (
      "bytes"
    	"encoding/binary"
    	"fmt"
    	"math"
    )
    
    func main() {
    	buf := new(bytes.Buffer)
    	var pi float64 = math.Pi
    
    	err := binary.Write(buf, binary.LittleEndian, pi)
    	if err != nil {
    		fmt.Println("binary.Write failed:", err)
    	}
    	fmt.Println(buf.Bytes()) //[24 45 68 84 251 33 9 64]
    }
