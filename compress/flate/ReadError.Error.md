# func (e *ReadError) Error() string

返回值：表示flate读取拷贝数据时的错误信息

功能说明：ReadError其实是一个struct，他实现了error接口，用于很方便的返回flate读取拷贝数据时的错误信息

示例：

flate读取拷贝数据时的错误信息，这里不再举例。