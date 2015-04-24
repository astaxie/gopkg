# func (e InternalError) Error() string

返回值：表示flate数据自身的错误信息

功能说明：InternalError其实是一个string，他实现了error接口，用于很方便的返回flate数据自身的错误信息

示例：

函数返回的error信息，这里不再举例。