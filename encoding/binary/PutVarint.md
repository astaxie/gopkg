## func PutVarint(buf []byte, x int64) int

###参数列表

- buf 需写入的缓冲区 
- x int64类型数字

###返回值：

- int 写入字节数。
- panic buf过小。

###功能说明：

PutVarint主要是将int64类型放入buf中，并返回写入的字节数。如果buf过小，PutVarint将抛出panic。

###代码实例：
    
    package main
    
    import (
      "encoding/binary"
    	"fmt"
    )
    
    func main() {
    	var i16 int16 = 1234
    	var i64 int64 = -1234567890
    	var sbuf []byte = make([]byte, 4)
    	var buf []byte = make([]byte, 10)
    
    	ret := binary.PutVarint(sbuf, int64(i16))
    	fmt.Println(ret, sbuf) // 2 [164 19 0 0]
    
    	//ret = binary.PutVarint(sbuf, i64) //panic
    
    	ret = binary.PutVarint(buf, i64)
    	fmt.Println(ret, buf) // 5 [163 139 176 153 9 0 0 0 0 0]
    }
