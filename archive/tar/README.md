# archive/tar包详解

常量列表：
// 类型
- TypeReg           = '0'    // 普通文件
- TypeRegA          = '\x00' // 普通文件
- TypeLink          = '1'    // 实体链接
- TypeSymlink       = '2'    // 指向链接
- TypeChar          = '3'    // 字符设备节点
- TypeBlock         = '4'    // 区块设备节点
- TypeDir           = '5'    // 目录
- TypeFifo          = '6'    // 命名管道节点
- TypeCont          = '7'    // 保留的
- TypeXHeader       = 'x'    // 扩展数据头
- TypeXGlobalHeader = 'g'    // 全局扩展数据头

变量列表：
- ErrWriteTooLong    = errors.New("archive/tar: write too long") // 写出太长
- ErrFieldTooLong    = errors.New("archive/tar: header field too long")	// 数据头太长
- ErrWriteAfterClose = errors.New("archive/tar: write after close")	// 在关闭写出流后操作
- ErrHeader = errors.New("archive/tar: invalid tar header")	// 无效的数据头

