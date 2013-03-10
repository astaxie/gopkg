# package fmt
## package 中的 format 说明

`import "fmt"`

Go 语言中的fmt包含有格式化输入输出的函数，类似于C语言的printf和scanf。格式字符串的规则来源于C但更简单更好用。

在 fmt 包中，有三种标准输出的打印语句：
> Printf 可以给定输出格式的打印语句
>
>*   Printf("格式化串定义",变量列表)
>
>*	格式化串定义：即format字符串
>
>*	变量列表: 可以是任意类型的值变量列表，系统实现的是一通用接口函数
>
> Print 这个函数不需要format字符串，等价于对每一个参数设置为%v
> 
> Println 等价于Print语句，但输出时会在参数之间加上空格并在输出结束后换行

	

### 1.例子 Exfmt01.go：

		// Exfmt01.go
		// %v  基本格式的值
		// %+v 当输出结构体时，扩展标志添加成员的名字。
		// %#v 值的Go语法表示。
		package main
		
		import "fmt"
		
		var timeZone = map[string]int{
		        "UTC": 0 * 60 * 60,
		        "EST": -5 * 60 * 60,
		        "CST": -6 * 60 * 60,
		        "MST": -7 * 60 * 60,
		        "PST": -8 * 60 * 60,
		}
		
		type T struct {
		        a int
		        b float32
		        c string
		}
		
		func main() {
		        t := &T{7, -2.35, "abc\tdef"}
		        fmt.Printf("No01:%v\n", t)
		        fmt.Printf("No02:%+v\n", t)
		        fmt.Printf("No03:%#v\n", t)
		        fmt.Printf("No04:%#v\n", timeZone)
		}


输出格式：

		No01:&{7 -2.35 abc   def}
        No02:&{a:7 b:-2.35 c:abc     def}
        No03:&main.T{a:7, b:-2.35, c:"abc\tdef"}
        No04:map[string]int{"MST":-25200, "UTC":0, "CST":-21600, "PST":-28800, "EST":-18000}


>重点说明：
>
>在go语言中，一般可以不考虑需打印数据的具体类型，直接利用格式中的 v 就可以了，方便啊！
>
>%v  基本格式的值
>
>%+v 当输出结构体时，扩展标志添加成员的名字。
>
>%#v 值的Go语法表示。

### 2.例子 Exfmt02.go：

        // Exfmt02.go
        // 测试和说明：
        // %t 值的true和false
        // %T 值的类型在Go语言中的表示
        // %% 打印一个 %
        // %q 给打印的字串自动加引号
        package main

        import "fmt"

        func main() {
                var Yes bool //bool 变量自动初始化为 false
                fmt.Printf("No01:%t\n", Yes)
                fmt.Printf("No02:%T\n", Yes)
                fmt.Printf("No03:%%\n")
                fmt.Printf("No04:%q\n","自动加引号")
        }
        // 输出：
        //No01:false
        //No02:bool
        //No03:%
        //No04:"自动加引号"

>重点说明：
>
>%t 用于打印bool类型的值，一般为true和false
>
>%T 值的类型在Go语言中的表示
>
>%% 打印一个 %
>

### 3.例子 Exfmt03.go：

        // Exfmt03.go
        // 测试和说明：
        // 整数部分
        // %b 二进制表示
        // %c 数值对应的Unicode编码字符
        // %d 十进制表示
        // %o 八进制表示
        // %x 十六进制表示，使用a-f
        // %X 十六进制表示，使用A-F
        // %U Unicode格式： U+1234，等价于"U+%04X"
        // %q 单引号
        package main

        import "fmt"

        func main() {
                fmt.Printf("No01:%b\n", 888)
                fmt.Printf("No02:%c\n", '太')
                fmt.Printf("No03:%d\n", 888)
                fmt.Printf("No04:%o\n", 888)
                fmt.Printf("No05:%x\n", 888)
                fmt.Printf("No06:%X\n", 888)
                fmt.Printf("No07:%U\n", 888)
                fmt.Printf("No08:%q\n", 88)
        }
        // 输出：
        //No01:1101111000
        //No02:太
        //No03:888
        //No04:1570
        //No05:378
        //No06:378
        //No07:U+0378
        //No08:'X'

>重点说明：
>
>整数部分
>
>%b 二进制表示
>
>%c 数值对应的Unicode编码字符
>
>%d 十进制表示
>
>%o 八进制表示
>
>%x 十六进制表示，使用a-f
>
>%X 十六进制表示，使用A-F
>
>%U Unicode格式： U+1234，等价于"U+%04X"
>
>%q 单引号
>

### 4.例子 Exfmt04.go：

        // Exfmt04.go
        // 测试和说明：
        // 浮点数：
        // %b 无小数部分、两位指数的科学计数法，和strconv.FormatFloat的'b'转换格式一致。
        // %e 科学计数法，举例：-1234.456e+78
        // %E 科学计数法，举例：-1234.456E+78
        // %f 有小数部分，但无指数部分，举例：123.456
        // %g 根据实际情况采用%e或%f格式（以获得更简洁的输出）
        // %G 根据实际情况采用%E或%f格式（以获得更简洁的输出）
        package main

        import "fmt"

        func main() {
                fmt.Printf("No01:%b\n", 888.66)
                fmt.Printf("No02:%e\n", 888.66)
                fmt.Printf("No03:%E\n", 888.66)
                fmt.Printf("No04:%f\n", 888.66)
                fmt.Printf("No05:%g\n", 888.33)
                fmt.Printf("No06:%g\n", 99999999.33)
                fmt.Printf("No07:%G\n", 888.33)
                fmt.Printf("No08:%G\n", 99999999.33)
        }
        // 输出：
        //No01:7816736025115361p-43
        //No02:8.886600e+02
        //No03:8.886600E+02
        //No04:888.660000
        //No05:888.33
        //No06:9.999999933e+07
        //No07:888.33
        //No08:9.999999933E+07

>重点说明：
>
>浮点数：
>
>%b 无小数部分、两位指数的科学计数法，和strconv.FormatFloat的'b'转换格式一致。
>
>%e 科学计数法，举例：-1234.456e+78
>
>%E 科学计数法，举例：-1234.456E+78
>
>%f 有小数部分，但无指数部分，举例：123.456
>
>%g 根据实际情况采用%e或%f格式（以获得更简洁的输出）
>
>%G 根据实际情况采用%E或%f格式（以获得更简洁的输出）
>

### 5.例子 Exfmt05.go：

        // Exfmt05.go
        // 测试和说明：
        // "No01:%.*d",10,123: 打印整数，并保证10为长度，关键是长度运行时动态传入
        // "No02:%6.*f",2,888.666: 打印浮点数，并保证两位小数，关键是小数位运行时动态传入
        package main

        import "fmt"

        func main() {
                fmt.Printf("No01:%.*d\n",10,123)
                fmt.Printf("No02:%6.*f\n",2,888.666)
        }
        // 输出：
        //No01:0000000123
        //No02:888.67

>重点说明：
>
>"No01:%.*d",10,123: 打印整数，并保证10为长度，关键是长度运行时动态传入
>
>"No02:%6.*f",2,888.666: 打印浮点数，并保证两位小数，关键是小数位运行时动态传入
>


***

## 字符串和byte切片类型说明：

	%s 直接输出字符串或者[]byte
	%q 双引号括起来的字符串
	%x 每个字节用两字符十六进制数表示（使用小写a-f）
	%X 每个字节用两字符十六进制数表示（使用大写A-F）
	
## 指针类型说明：

	%p 0x开头的十六进制数表示
	木有'u'标志。如果是无类型整数，自然会打印无类型格式。类似的，没有必要去区分操作数的大小(int8, int64)。
	宽度和精度格式化控制是指的Unicode编码字符的数量（不同于C的printf，它的这两个因子指的是字节的数量。）两者均可以使用'*'号取代（任一个或两个都），此时它们的值将被紧接着的参数控制，这个操作数必须是整型。
	对于数字，宽度设置总长度，精度设置小数部分长度。例如，格式%6.2f 输出123.45。
	对于字符串，宽度是输出字符数目的最低数量，如果不足会用空格填充。精度是输出字符数目的最大数量，超过则会截断。
	
## 其他符号说明：	

	+ 总是输出数值的正负号；对%q(%+q)将保证纯ASCII码输出
	- 用空格在右侧填充空缺而不是默认的左侧。
	# 切换格式：在八进制前加0(%#o)，十六进制前加0x(%#x)或0X(%#X)；废除指针的0x(%#p)；
	对%q (%#q)如果可能的话输出一个无修饰的字符串；
	对%U(%#U)如果对应数值是可打印字符输出该字符。
	' ' 对数字(% d)空格会留一个空格在数字前并忽略数字的正负号；
	对切片和字符串(% x, % X)会以16进制输出。
	0 用前置0代替空格填补空缺。

## 补充说明：

	每一个类似Printf的函数，都会有一个同样的Print函数，此函数不需要format字符串，等价于对每一个参数设置为%v。另一个变体Println会在参数之间加上空格并在输出结束后换行。
	如果参数是一个接口值，将使用内在的具体实现的值，而不是接口本身，%v参数不会被使用。如下：
	var i interface{} = 23
	fmt.Printf("%v\n", i)
	将输出23。
	如果参数实现了Formatter接口，该接口可用来更好的控制格式化。
	如果格式（标志对Println等是隐含的%v）是专用于字符串的(%s %q %v %x %X)，还提供了如下两个规则：
	1. 如果一个参数实现了error接口，Error方法会用来将目标转化为字符串，随后将被按给出的要求格式化。
	2. 如果参数提供了String方法，这个方法将被用来将目标转换为字符串，然后将按给出的格式标志格式化。
	为了避免有可能的递归循环，例如：
	type X string
	func (x X) String() string { return Sprintf("<%s>", x) }
	会在递归循环前转换值：
	func (x X) String() string { return Sprintf("<%s>", string(x)) }
	错误的格式：
	如果提供了一个错误的格式标志，例如给一个字符串提供了%d标志，生成的字符串将包含对该问题的描述，如下面的例子：
	错误或未知的格式标志: %!verb(type=value)
	Printf("%d", hi): %!d(string=hi)
	太多参数: %!(EXTRA type=value)
	Printf("hi", "guys"): hi%!(EXTRA string=guys)
	缺少参数: %!verb(MISSING)
	Printf("hi%d"): hi %!d(MISSING)
	使用非整数提供宽度和精度: %!(BADWIDTH) or %!(BADPREC)
	Printf("%*s", 4.5, "hi"): %!(BADWIDTH)hi
	Printf("%.*s", 4.5, "hi"): %!(BADPREC)hi
	所有的错误都使用"%!"起始，（紧随单字符的格式标志）以括号包围的错误描述结束。
	输入
	一系列类似的函数读取格式化的文本，生成值。Scan，Scanf和Scanln从os.Stdin读取；Fscan，Fscanf和Fscanln 从特定的io.Reader读取；Sscan，Sscanf和Sscanln 从字符串读取；Scanln，Fscanln和Sscanln在换行时结束读取，并要求数据连续出现；Scanf，Fscanf和Sscanf会读取一整行以匹配格式字符串；其他的函数将换行看着空格。
	Scanf, Fscanf, and Sscanf根据格式字符串解析数据，类似于Printf。例如，%x将读取一个十六进制数，%v将读取值的默认表示。
	格式行为类似于Printf，但有如下例外：
	%p没有提供
	%T没有提供
	%e %E %f %F %g %G是等价的，都可以读取任何浮点数或者复合数（非复数，指科学计数法表示的带指数的数）
	%s 和 %v字符串使用这两个格式读取时会因为空格而结束
	不设格式或者使用%v读取整数时，如果前缀为0(八进制)或0x(十六进制)，将按对应进制读取。
	宽度在输入中被解释（%5s意思是最多从输入读取5个字符赋值给一个字符串），但输入系列函数没有解释精度的语法(木有%5.2f，只有%5f)。
	输入系列函数中的格式字符串，所有非空的空白字符（除了换行符之外），无论在输入里还是格式字符串里，都等价于1个空白字符。格式字符串必须匹配输入的文本，如果不匹配将停止读取数据并返回函数已经赋值的参数的数量。
	所有的scan系列函数，如果参数包含Scan方法（或者说实现了Scanner接口），该参数将使用该方法读取文本。另外，如果被填写的参数的数量少于提供的参数的数量，将返回一个错误。
	所有要被输入的参数都应该是基础类型或者实现了Scanner接口的数据类型的指针。
	注意：Fscan等函数可以从输入略过一些字符读取需要的字符并返回，这就意味着一个循环的读取程序可能会跳过输入的部分数据。当数据间没有空白时就会导致出现问题。如果读取这提供给Fscan系列函数ReadRune 方法，这个方法可以用来读取字符。如果读取者还提供了UnreadRune 方法，该方法将被用来保存字符以使成功的调用不会丢失数据。为了给一个没有这些功能的读取者添加这俩方法，使用bufio.NewReader。

# 函数及相关类型的简介

## 函数说明

	func Errorf
	func Errorf(format string, a ...interface{}) error
	Errorf根据格式字符串和参数表生成一个字符串，该字符串满足error接口。
	func Fprint
	func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	Fprint将所有参数都使用默认的格式写入w。如果相邻两个参数都不是字符串时，会在参数间添加空白。函数返回写入的字节数和任何遇到的错误。
	func Fprintf
	func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	Fprintf根据格式字符串将参数写入w。函数返回写入的字节数和任何遇到的错误。
	func Fprintln
	func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	Fprintln将所有参数都使用默认的格式写入w并在最后添加换行。如果相邻两个参数都不是字符串时，会在参数间添加空白。函数返回写入的字节数和任何遇到的错误。
	func Fscan
	func Fscan(r io.Reader, a ...interface{}) (n int, err error)
	Fscan从r读取文本，将连续的空白分割的数据存入连续的参数里。换行视同空白。它返回成功读取的参数的数量。如果少于提供的参数的数量，返回值err将报告原因。
	func Fscanf
	func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
	Fscanf从r读取文本，根据格式字符串顺序将数据存入参数中。它返回成功解析并存入的参数的数量。
	func Fscanln
	func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
	Fscanln类似Fscanf，但会在换行符中止，并且存入最后一条后时读取位置必须有换行或者结束符。
	func Print
	func Print(a ...interface{}) (n int, err error)
	Print将所有参数都使用默认的格式写入标准输出。如果相邻两个参数都不是字符串时，会在参数间添加空白。函数返回写入的字节数和任何遇到的错误。
	func Printf
	func Printf(format string, a ...interface{}) (n int, err error)
	Printf根据格式字符串将参数写入标准输出。函数返回写入的字节数和任何遇到的错误。
	func Println
	func Println(a ...interface{}) (n int, err error)
	Println将所有参数都使用默认的格式写入标准输出并在最后添加换行。如果相邻两个参数均非字符串时，会在参数间添加空白。函数返回写入的字节数和任何遇到的错误。
	func Scan
	func Scan(a ...interface{}) (n int, err error)
	Scan从标准输入读取文本，将空白分割的连续数据顺序存入参数里。换行视同空白。它返回成功读取的参数的数量。如果少于提供的参数的数量，返回值err将报告原因。
	func Scanf
	func Scanf(format string, a ...interface{}) (n int, err error)
	Scanf从标准输入读取文本，根据格式字符串顺序将数据存入参数中。它返回成功解析并存入的参数的数量。
	func Scanln
	func Scanln(a ...interface{}) (n int, err error)
	Scanln类似Scan，但会在换行符中止，并且存入最后一条后时读取位置必须有换行或者结束符。
	func Sprint
	func Sprint(a ...interface{}) string
	Sprint将所有参数都使用默认的格式写入并生成一个字符串。如果相邻两个参数都不是字符串时，会在参数间添加空白。
	func Sprintf
	func Sprintf(format string, a ...interface{}) string
	Sprintf根据格式字符串将参数写入并返回生成的字符串。
	func Sprintln
	func Sprintln(a ...interface{}) string
	Sprintln将所有参数都使用默认的格式写入并生成一个字符串。如果相邻两个参数都不是字符串时，会在参数间添加空白。字符串最后会添加换行符。
	func Sscan
	func Sscan(str string, a ...interface{}) (n int, err error)
	Sscan从字符串读取文本，将空白分割的连续数据顺序存入参数里。换行视同空白。它返回成功读取的参数的数量。如果少于提供的参数的数量，返回值err将报告原因。
	func Sscanf
	func Sscanf(str string, format string, a ...interface{}) (n int, err error)
	Scanf从字符串读取文本，根据格式字符串顺序将数据存入参数中。它返回成功解析并存入的参数的数量。
	func Sscanln
	func Sscanln(str string, a ...interface{}) (n int, err error)
	Sscanln类似Sscan，但会在换行符中止，并且存入最后一条后时读取位置必须有换行或者结束符。

## 类型说明
	
	type Formatter
	type Formatter interface {
	Format(f State, c rune)
	}
	Formatter是一个供用户定制的格式化接口。Format方法的实现可能需要调用Sprintf 或Fprintf(f)等函数来生成输出。
	type GoStringer
	type GoStringer interface {
	GoString() string
	}
	GoStringer接口由任意包含GoString方法的数据实现，这个方法定义了数据的Go语法格式。GoString方法用来在使用%#v格式标志时输出值。
	type ScanState
	type ScanState interface {
	// ReadRune函数从输入读取下一个Unicode符号。如果在Scanln，Fscanln或Sscanln中调用，本函数会在读取到第一个'\n'或达到最大宽度时返回EOF。
	ReadRune() (r rune, size int, err error)
	// UnreadRune会让ReadRune的下一次调用返回同一个字符。
	UnreadRune() error
	// SkipSpace跳过输入的空白。换行被视为空白（Scanln，Fscanln和Sscanln例外，这三个函数里换行符视为EOF）。
	SkipSpace()
	// Token方法会在skipSpace为真时跳过输入中的空白，并返回一个满足f(c)的Unicode字符。如果f是nil，则使用!unicode.IsSpace(c)（即返回第一个非空格Unicode字符）；
	// 即是说，本函数只对非空字符起效。换行符视为空白字符（Scanln，Fscanln和Sscanln例外，这三个函数里换行符视为EOF）。
	// 返回的切片类型指向共享的数据，该数据可以被下一次Token的调用（使用ScanState接口作为输入调用Scan函数）中或者调用返回的Scan方法时重写。
	Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
	// Width返回width选项的值以及其是否被设定。
	Width() (wid int, ok bool)
	// 因为ReadRune用接口实现，Read方法应该永远不被scan程序调用，一个好使的ScanState实现应该保证总是从Read返回错误。
	Read(buf []byte) (n int, err error)
	}
	ScanState是一个交给用户定制的Scanner接口的参数的接口。Scanner接口可能会进行一次一个字符的扫描或者要求ScanState去探测下一个空白分隔的token。
	type Scanner
	type Scanner interface {
	Scan(state ScanState, verb rune) error
	}
	任何实现了Scan方法的对象都实现了Scanner接口，Scan方法会从输入读取数据并将处理结果存入接受端，接收端必须是有效的指针。Scan方法会被任何Scan、Scanf、Scanln等函数调用，只要对应的参数实现了该方法。
	type State
	type State interface {
	// Write可被调用以发出格式化的输出。
	Write(b []byte) (ret int, err error)
	// Width返回宽度的值及其是否被设定。
	Width() (wid int, ok bool)
	// Precision返回精度的值及其是否被设定。
	Precision() (prec int, ok bool)
	// Flag返回符号（正负号……）的值是否被设定。
	Flag(c int) bool
	}
	State是一个提供给Formatter接口的输出参数的接口。它提供对io.Writer接口的、使用提供的数据对参数格式化后的访问。
	type Stringer
	type Stringer interface {
	String() string
	}
	Stringer接口被任何实现了String方法的类型自动实现，该方法定义了该类型的“原生”格式。String方法用来输出参数，当使用%s或%v格式时，或者被Print等不使用格式字符串的函数输出时。 
	
