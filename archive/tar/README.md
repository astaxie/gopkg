# archive/tar包详解

### 常量列表
	const(
		TypeReg           = '0'    // 普通文件
		TypeRegA          = '\x00' // 普通文件
		TypeLink          = '1'    // 实体链接
		TypeSymlink       = '2'    // 指向链接
		TypeChar          = '3'    // 字符设备节点
		TypeBlock         = '4'    // 区块设备节点
		TypeDir           = '5'    // 目录
		TypeFifo          = '6'    // 命名管道节点
		TypeCont          = '7'    // 保留的
		TypeXHeader       = 'x'    // 扩展数据头
		TypeXGlobalHeader = 'g'    // 全局扩展数据头
	)
### 变量列表：
	var(
		ErrWriteTooLong    = errors.New("archive/tar: write too long") // 写出太长
		ErrFieldTooLong    = errors.New("archive/tar: header field too long")	// 数据头太长
		ErrWriteAfterClose = errors.New("archive/tar: write after close")	// 在关闭写出流后操作
		ErrHeader = errors.New("archive/tar: invalid tar header")	// 无效的数据头
	)

### 数据头结构
	type Header struct {
    	Name       string    // 写入数据头的文件名称
		Mode       int64     // 许可与模式比特位(bits)
    	Uid        int       // 拥有者的用户ID
    	Gid        int       // 拥有者的群组ID
    	Size       int64     // 以字节(bytes)为单位的长度
    	ModTime    time.Time // 修改时间
    	Typeflag   byte      // 写入的数据头类型（文件类型）
    	Linkname   string    // 链接的目标名称
    	Uname      string    // 拥有者的用户名
    	Gname      string    // 拥有者的群组名
    	Devmajor   int64     // 字符或区块设备的主要数字
    	Devminor   int64     // 字符或区块设备的次要数字
    	AccessTime time.Time // 存取时间
    	ChangeTime time.Time // 权限状态改变时间
	}

*	[Reader结构 - 读取tar文件](Reader.md) 
*	[Writer结构 - 创建tar文件](Writer.md) 