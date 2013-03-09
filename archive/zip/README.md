# archive/zip包详解

###常量列表
	// 压缩的方法
	const (
    	Store   uint16 = 0
    	Deflate uint16 = 8
	)
###变量列表：
	var (
    	ErrFormat    = errors.New("zip: not a valid zip file")	// 不是一个有效的zip文件
    	ErrAlgorithm = errors.New("zip: unsupported compression algorithm")	// 不支持的压缩算法
    	ErrChecksum  = errors.New("zip: checksum error")	// 和校验错误
	)
###文件结构
	type File struct {
 		FileHeader	// 文件数据头
    	// 包含其它被过滤或未输出的字段
	}
###文件数据头结构
	type FileHeader struct {
    	Name             string // 写入数据头的文件名称
    	CreatorVersion   uint16 // 创建者版本
    	ReaderVersion    uint16 // 读取者版本
    	Flags            uint16 // 写入标志
    	Method           uint16 // 写入方法
    	ModifiedTime     uint16 // MS-DOS 时间
    	ModifiedDate     uint16 // MS-DOS 日期
    	CRC32            uint32 // CRC32校验和
    	CompressedSize   uint32 // 压缩后体积
    	UncompressedSize uint32 // 未压缩体积
    	Extra            []byte // 附加属性
    	ExternalAttrs    uint32 // 外部属性（依赖于创建者版本）
    	Comment          string // 文件注释
}

*	[Reader结构 － 读取zip文件](Reader.md)
*	[Writer结构 － 创建zip文件](Writer.md)