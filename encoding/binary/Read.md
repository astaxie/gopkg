## func Read(r io.Reader, order ByteOrder, data interface{}) error

###参数列表

- buf 需写入的缓冲区 
- x uint64类型数字

###返回值：

- int 写入字节数。
- panic buf过小。

###功能说明：

PutUvarint主要是将uint64类型放入buf中，并返回写入的字节数。如果buf过小，putUvarint将抛出panic。

###代码实例：
