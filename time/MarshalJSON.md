# func (t Time) MarshalJSON() ([]byte, error)

参数列表：

- 无

返回值：

- []byte json串
- error 错误

功能说明：

MarshalJSON实现了json.Marshaler接口。时间t按照RFC3339格式化。

代码实例：

无