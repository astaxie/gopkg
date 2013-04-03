# compress/gzip包详解

###常量列表
	// 压缩的方法
	const (
    	NoCompression      = flate.NoCompression	// 不压缩
    	BestSpeed          = flate.BestSpeed	// 最快速度
    	BestCompression    = flate.BestCompression	// 最佳压缩比
    	DefaultCompression = flate.DefaultCompression	// 默认压缩比
	)
###变量列表：
	var (
		// 当读取gzip数据时发现无效的校验和时将返回该错误
    ErrChecksum = errors.New("gzip: invalid checksum")
    	// 当读取gzip数据时发现无效的数据头时将返回该错误
    ErrHeader = errors.New("gzip: invalid header")
	)
###数据头结构
	type Header struct {
    	Comment string    // 文件注释
    	Extra   []byte    // 附加数据
    	ModTime time.Time // 文件修改时间
    	Name    string    // 文件名
    	OS      byte      // 操作系统类型
	}
	
*	[Reader结构 － 读取gzip文件](Reader.md)
*	[Writer结构 － 创建gzip文件](Writer.md)