# Method(int) Method

参数列表

- int 传入方法的索引号

返回值：

- .Method  返回函数的 reflect.Method 类型

功能说明：

- reflect.TypeOf(x).Method(x) 指定索引号返回struct结构绑定的方法

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 int
	}
	func (f A) test(){}
	func (f A) test1(){}
	func (f A) test2(){}
	
	func main(){
		var a A
		var typeof reflect.Type = reflect.TypeOf(a)
		for i:=0; i<typeof.NumMethod(); i++ {
			m := typeof.Method(i)
			fmt.Println(m.Index, m.PkgPath, m.Name, m.Type, m.Func)
			// 1 >>0 main test func(main.A) <func(main.A) Value>
			// 2 >>1 main test1 func(main.A) <func(main.A) Value>
			// 3 >>2 main test2 func(main.A) <func(main.A) Value>
		}
	}
