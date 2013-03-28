# func DeepEqual(a1, a2 interface{}) bool

参数列表

- a1 interface{} 比较A
- a2 interface{} 比较B

返回值：

- bool 如果两个参数的类型与值是相等的返回true,否则返回false。

功能说明：

- DeepEqual 深刻的平等判断。在可能的情况下它使用 默认==平等，将扫描Array，Slice，Map，Struct和字段（Field）的成员。正确处理递归类型。如果他们俩都nil，那么函数是相等的。

代码实例1：

  	package main
		import (
		    "fmt"
		    "reflect"
		)
		func main() {
			var a int
			var b int
			a = 1
			b = 2
			var booL = reflect.DeepEqual(a, b)
			fmt.Println(booL)
			//>>false
			b = 1
			booL = reflect.DeepEqual(a, b)
			fmt.Println(booL)
			//>>true
		}

代码实例2：

		package main
		import (
		    "fmt"
		    "reflect"
		)
		func main() {
		 	type A struct {
				A0 int
			}
			type B struct {
				A0 int
			}
			var a A
			var b B
			var booL = reflect.DeepEqual(a, b)
			fmt.Println(booL)
			//>>false

			booL = reflect.DeepEqual(a.A0, b.A0)
			fmt.Println(booL)
			//>>true
		}

