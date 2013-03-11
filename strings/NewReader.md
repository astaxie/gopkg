# type Reader
 Reader是通过读取一个字符串之后实现了io.Reader, io.ReaderAt, io.Seeker, io.ByteScanner, 和io.RuneScanner 接口

# func NewReader(s string) *Reader
参数列表

- s 读取的字符串

返回值：

- *Reader 通过读取一个字符串之后返回Reader对象

对象的方法列表：

- func (r *Reader) Len() int   //返回未读取的字符串的长度
- func (r *Reader) Read(b []byte) (n int, err error)  //读取数据到b中，返回读取的实际大小n，如果出错返回err，例如EOF或者b的长度为0
- func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) //按照指定的off位置开始读取内容到b，返回读取的实际大小n，如果出错返回err，例如off小于0或者大于本身的长度或者文件尾
- func (r *Reader) ReadByte() (b byte, err error)  //读取一个byte的数据
- func (r *Reader) ReadRune() (ch rune, size int, err error)  //读取一个rune的数据
- func (r *Reader) Seek(offset int64, whence int) (int64, error) //根据whence来移动offset，如果whence=0为直接移动offset位置，=1为移动到当前位置之后的offset，=2为移动到当前字符串长度之后的offset位置
- func (r *Reader) UnreadByte() error  //当前读取的位置向前移一个byte
- func (r *Reader) UnreadRune() error  //当前读取的位置向前移一个rune

功能说明：

该函数主要是通过把字符串读取Reader之后进行的一些读取操作

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		read := strings.NewReader("I am asta谢")
		var b []byte
		fmt.Println(read.Len())  //12
		b = make([]byte,8)
		n, err := read.Read(b)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(b)   //[73 32 97 109 32 97 115 116]
		fmt.Println(n)   //8
		n, err = read.ReadAt(b,3)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(b)   //[109 32 97 115 116 97 232 176]
		fmt.Println(n)   //8
		bt,err:=read.ReadByte()	
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(bt)  //97
		rn,size,err:=read.ReadRune()	
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(rn)    //35874
		fmt.Println(size)  //3
	}
