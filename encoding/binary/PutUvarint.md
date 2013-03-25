## func PutUvarint(buf []byte, x uint64) int

###参数列表

- buf 需写入的缓冲区 
- x uint64类型数字

###返回值：

- int 写入字节数。
- panic buf过小。

###功能说明：

PutUvarint主要是将uint64类型放入buf中，并返回写入的字节数。如果buf过小，putUvarint将抛出panic。

###代码实例：
    
    package main
    
    import (
        "encoding/binary"
    	"fmt"
    )
    
    func main() {
    	var u16 uint16 = 1234 //0x1020
    	var u64 uint64 = 0x1020304040302010
    	var sbuf []byte = make([]byte, 4)
    	var buf []byte = make([]byte, 10)
    
    	ret := binary.PutUvarint(sbuf, uint64(u16))
    	fmt.Println(ret, sbuf) // 2 [210 9 0 0]
    
    	//ret = binary.PutUvarint(sbuf, u64) //panic
    
    	ret = binary.PutUvarint(buf, u64)
    	fmt.Println(ret, buf) //9 [144 192 192 129 132 136 140 144 16 0]
    }
