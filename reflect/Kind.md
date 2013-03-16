# type Kind

参数列表

- 无

返回值：

- 无

功能说明：

- 一个类型代表一种特定类型的类型，零Kind不是有效的Kind。
    - reflect.Kind 有以下常量成员
		- reflect.Invalid&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 无效
		- reflect.Bool&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 布尔
		- reflect.Int&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 整数（有符号）
		- reflect.Int8&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 整数8位（有符号）
		- reflect.Int16&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 整数16位（有符号）
		- reflect.Int32&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 整数32位（有符号）
		- reflect.Int64&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 整数64位（有符号）
		- reflect.Uint&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 整数（无符号）
		- reflect.Uint8&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 整数8（无符号）
		- reflect.Uint16&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 整数16（无符号）
		- reflect.Uint32&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 整数（无符号）
		- reflect.Uint64&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 整数（无符号）
		- reflect.Uintptr&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 整数（指针,无符号）
		- reflect.Float32&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 浮点数32位
		- reflect.Float64&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 浮点数64位
		- reflect.Complex64&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 复数64位
		- reflect.Complex128&nbsp;&nbsp;&nbsp;&nbsp;// 复数128位
		- reflect.Array&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 数组
		- reflect.Chan&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 信道
		- reflect.Func&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 函数
		- reflect.Interface&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 接口
		- reflect.Map&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 地图
		- reflect.Ptr&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 指针
		- reflect.Slice&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 切片
		- reflect.String&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 字符
		- reflect.Struct&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;// 结构
		- reflect.UnsafePointer&nbsp;// 安全指针

代码实例：

		package main
		import (
		    "fmt"
		    "reflect"
		)
		func main() {
			var a string
			var kind reflect.Kind = reflect.TypeOf(a).Kind()
			fmt.Println(kind, kind == reflect.String, kind == reflect.Int)
			//>>string true false
		}
