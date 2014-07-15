# func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte

参数列表

- dst     原列表
- f       需要append到列表的浮点数
- fmt     'f'为小数点格式，'e'为科学计数法
- prec    表示小数点后位数
- bitSize 表示

返回值:

- []byte  返回列表

功能说明：

- 将浮点数f转换为字符串值，并将转换结果追加到dst的尾部，返回追加后的[]byte。fmt为显示格式。prec是精确的位数，bitSize是64或者32

代码示例：

	package main
	
	import (
		"fmt"
		"strconv"
	)
	
	func main() {
		f := 100.123456789
		fmt.Println(f)
		c := strconv.AppendFloat(make([]byte, 0), f, 'f', 10, 32)
		fmt.Println(string(c))
		c = strconv.AppendFloat(make([]byte, 0), f, 'e', 10, 32)
		fmt.Println(string(c))
		c = strconv.AppendFloat(make([]byte, 0), f, 'f', 10, 64)
		fmt.Println(string(c))
		c = strconv.AppendFloat(make([]byte, 0), f, 'e', 10, 64)
		fmt.Println(string(c))
	}


代码输出：

	100.123456789
	100.1234588623
	1.0012345886e+02
	100.1234567890
	1.0012345679e+02
