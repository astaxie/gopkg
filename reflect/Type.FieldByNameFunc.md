# FieldByNameFunc(match func(string) bool) (StructField, bool)

参数列表

- match func(string) bool 传入一个函数用于筛选匹配的字段。函数又传入字段的“字符串”名称并判断，func 返回 true，表示找到匹配。

返回值：

- .StructField 返回字段的 reflect.StructField 类型
- bool 如果找到字段，返回true，则返回false

功能说明：

- reflect.TypeOf(interface{}).FieldByNameFunc(func) 筛选并判断满足匹配的字段。如果找到了该字段的名称，返回的第一个结构字段和一个布尔值。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 int "标签0"
		A1 string "标签1"
		A2 []byte "标签2"
		A3 *A "指针指向自身"
	}

	func main(){
		var a A
		var typeof reflect.Type = reflect.TypeOf(a)

		var fun = func(s string) bool { //遍历字段，传入每个字段的字段名
			if s == "A3" { //请结合正则匹配，才能发挥最大效用。
				return true
			}
			return false
		}
		var f, ok = typeof.FieldByNameFunc(fun)
		fmt.Println(ok, f.Index, f.PkgPath, f.Name, f.Type, f.Tag, f.Anonymous)
		//>>true [3]  A3 *main.A 指针指向自身 false
	}
