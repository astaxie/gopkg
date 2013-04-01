## func Uvarint(buf []byte) (uint64, int)

###参数列表

- buf 需要解码的缓冲区 

###返回值：

- uint64 解码的数据。
- int 解析的字节数。

###功能说明：

Uvarint是从buf中解码并返回一个uint64的数据，及解码的字节数（>0）。
如果出错，则返回数据0和一个小于等于0的字节数n，其意义为：

    n == 0：buf太小
    n  < 0：数据太大，超过64位，且-n为已解析字节数

###代码实例：

    package main
    
    import (
      "encoding/binary"
    	"fmt"
    )
    
    func main() {
    	var sbuf []byte
    	var buf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1}
    	var bbuf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}
    
    	num, ret := binary.Uvarint(sbuf)
    	fmt.Println(num, ret) //0 0
    
    	num, ret = binary.Uvarint(buf)
    	fmt.Println(num, ret) //1161981756374523920 9
    
    	num, ret = binary.Uvarint(bbuf)
    	fmt.Println(num, ret) //0 -11
    }
