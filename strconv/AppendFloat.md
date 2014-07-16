# func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte

参数列表

- dst     原列表
- f       需要append到列表的浮点数
- fmt     转换格式 'b' 'e' 'E' 'f' 'g'或'G'
- prec    浮点数精度
- bitSize 32或64，32对应float32，64对应float64

返回值:

- []byte  返回列表

功能说明：

- 将浮点数f转换为字符串值，并将转换结果追加到dst的尾部，返回追加后的[]byte。
- 浮点数格式有'b' (-ddddp±ddd, 二进制指数), 'e' (-d.dddde±dd, 十进制指数), 'E' (-d.ddddE±dd, 十进制指数), 'f' (-ddd.dddd, 无指数), 'g' (大指数时相当于'e', 其他情况时相当于'f'), 'G' (大指数时相当于'E', 其他情况相当于'f').
- 精度用于控制当格式为'e' 'E' 'f' 'g' 'G'时除指数外的数字的个数；对于'e' 'E' 'f'指小数点后位数；对于'g' 'G'则表示总共的位数；如果使用-1，表示不改变数值的最小位数

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
