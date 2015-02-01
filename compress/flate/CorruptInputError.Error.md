# func (e CorruptInputError) Error() string

返回值：在输入的指定偏移量位置存在损坏的错误信息

功能说明：CorruptInputError其实是一个int64，他实现了error接口，用于很方便的返回输入的指定偏移量位置存在损坏的错误信息

示例：

可能是一些读取、写入、拷贝等函数返回的error信息，这里不再举例。