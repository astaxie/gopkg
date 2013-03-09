	const (
    	ScanIdents     = 1 << -Ident
    	ScanInts       = 1 << -Int
    	ScanFloats     = 1 << -Float // includes Ints
    	ScanChars      = 1 << -Char
    	ScanStrings    = 1 << -String
    	ScanRawStrings = 1 << -RawString
    	ScanComments   = 1 << -Comment
    	SkipComments   = 1 << -skipComment // if set with ScanComments, comments become white space
    	GoTokens       = ScanIdents | ScanFloats | ScanChars | ScanStrings | ScanRawStrings | ScanComments | SkipComments
	)

用于控制词法单元识别的预定义模式位。比如，通过以下模式位设置，可以配置一个只识别（Go）标识符、整型、并且忽略注释的词法解析器：
	
	ScanIdents | ScanInts | SkipComments

	const (
    	EOF = -(iota + 1)
    	Ident
    	Int
    	Float
    	Char
    	String
    	RawString
    	Comment
	)

通过Scan方法可以获取的词法单元，此外还可以获取Unicode字符。
	
	const GoWhitespace = 1<<'\t' | 1<<'\n' | 1<<'\r' | 1<<' '

GoWhitespace，Scanner的空白字符的默认值。

