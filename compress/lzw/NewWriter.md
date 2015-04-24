# func NewWriter(w io.Writer, order Order, litWidth int) io.WriteCloser

参数列表：

- w 输出压缩数据的io.Writer  
- order lzw数据流的位顺序，有LSB和MSB
- litWidth 字面码的位数，必须在[2,8]范围内，一般为8

返回值：一个io.WriteCloser，可以将压缩的数据写入其下层的w，调用者使用后要将其关闭

功能说明:

NewWriter创建一个新的io.WriteCloser。它将数据压缩后写入w。调用者要在写入结束之后调用返回io.WriteCloser的Close函数关闭。litWidth是字面码的为数，必须在[2,8]范围内，一般为8.

示例：

	package main
	
	import (
		"bytes"
		"compress/lzw"
		"fmt"
		"io"
		"os"
	)
	
	func main() {
		//一个缓冲区存储压缩的内容
		buf := bytes.NewBuffer(nil)
	
		w := lzw.NewWriter(buf, lzw.LSB, 8)
		//写入待压缩内容
		w.Write([]byte("compress/flate\n"))
		w.Close()
		fmt.Println(buf)
	
		//解压
		r := lzw.NewReader(buf, lzw.LSB, 8)
		defer r.Close()
		io.Copy(os.Stdout, r)
	}