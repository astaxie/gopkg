## func ReadVarint(r io.ByteReader) (int64, error)

###参数列表

- r 实现ReadByte()接口的对象

###返回值：

- int64 解析出的数据
- error 返回的错误

###功能说明：

ReadVarint从r中解析并返回一个int64类型的数据及出现的错误。

###代码实例：

    package main
    
    import (
      "bytes"
    	"encoding/binary"
    	"fmt"
    )
    
    func main() {
    	var sbuf []byte
    	var buf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1}
    	var bbuf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}
    
    	num, err := binary.ReadVarint(bytes.NewBuffer(sbuf))
    	fmt.Println(num, err) //0 EOF
    
    	num, err = binary.ReadVarint(bytes.NewBuffer(buf))
    	fmt.Println(num, err) //580990878187261960 <nil>
    
    	num, err = binary.ReadVarint(bytes.NewBuffer(bbuf))
    	fmt.Println(num, err) //2310373135097532424 binary: varint overflows a 64-bit integer
    }
