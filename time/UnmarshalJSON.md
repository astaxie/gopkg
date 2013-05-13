# func (t *Time) UnmarshalJSON(data []byte) (err error)

参数列表：

- data json数据

返回值：

- err 错误

功能说明：

UnmarshalJSON实现了json.Unmarshaler接口。t按照RFC3339格式化。

代码实例：

无