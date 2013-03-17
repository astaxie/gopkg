# type Method

参数列表

- 无

返回值：

- .Index 返回当前方法的索引号
- .Name 返回当前方法的名称
- .PkgPath 返回当前方法包的路径
- .Type 返回当前方法的 Type 类型
- .Func 返回当前方法的 Value 类型

功能说明：

- Method 是一种单独的方法。

代码实例：

  	package main
		import (
		    "fmt"
		    "reflect"
		)
		
		type A struct {
		}
		func (A) test(){}
		
		func main(){
			var a A
			var method reflect.Method = reflect.TypeOf(a).Method(0) //0 表示是 a 结构中的第几位函数
			fmt.Println(method.Index, method.Name, method.PkgPath, method.Type, method.Func)
			//>>0 test main func(main.A) <func(main.A) Value>
		}
