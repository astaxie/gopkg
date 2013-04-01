## func Read(r io.Reader, order ByteOrder, data interface{}) error

###参数列表

- r 可以读出字节流的数据源 
- order 特殊字节序，包中提供大端字节序和小段字节序
- data 需要解码成的数据

###返回值：

- error 返回错误

###功能说明：

Read从r中读出字节数据并反序列化成结构数据。data必须是固定长的数据值或固定长数据的slice。从r中读出的数据可以使用特殊的
字节序来解码，并顺序写入value的字段。当填充结构体时，使用(_)名的字段将被跳过。例如：为了字节对齐而添加的空字段。

###代码实例：
    package main
    
    import (
      "bytes"
    	"encoding/binary"
    	"fmt"
    )
    
    func main() {
    	var pi float64
    	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
    	buf := bytes.NewBuffer(b)
    
    	err := binary.Read(buf, binary.LittleEndian, &pi)
    	if err != nil {
    		fmt.Println("binary.Read failed:", err)
    	}
    	fmt.Println(pi) //3.141592653589793
    }
