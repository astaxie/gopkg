## func ReadUvarint(r io.ByteReader) (uint64, error)

###参数列表

- r 实现ReadByte()接口的对象

###返回值：

- uint64 解析出的数据
- error 返回的错误

###功能说明：

ReadUvarint从r中解析并返回一个uint64类型的数据及出现的错误。

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
    
    	num, err := binary.ReadUvarint(bytes.NewBuffer(sbuf))
    	fmt.Println(num, err) //0 EOF
    
    	num, err = binary.ReadUvarint(bytes.NewBuffer(buf))
    	fmt.Println(num, err) //1161981756374523920 <nil>
    
    	num, err = binary.ReadUvarint(bytes.NewBuffer(bbuf))
    	fmt.Println(num, err) //4620746270195064848 binary: varint overflows a 64-bit integer
    }
