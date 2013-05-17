# type SelectCase struct
参数列表

- 无

返回值：

- 无

功能说明：

- 预前设置Select目的，结合reflect.Select函数使用。若设置 Dir 值为SelectDefault，Chan与Send无需设置。Chan表示接收。Chan与Send不可同时现设置。
- 要懂这个类型之前，请先看 reflect.Select 函数如何使用，这样会很快就明白下面代码实例是什么意思。

  type SelectCase struct {
		    Dir  SelectDir 	// select方向，SelectDir 的值就是 SelectSend、SelectRecv、SelectDefault
		    Chan Value		// 通道使用（接收）
		    Send Value		// 发送的值（发送）
		}

代码实例1：

    package main
    import (
        "fmt"
        "reflect"
    )
    
    func main(){
		var chs = make(chan int, 1)
		var selectCase = make([]reflect.SelectCase, 1)
		selectCase[0].Dir = reflect.SelectRecv //2
		selectCase[0].Chan = reflect.ValueOf(chs) //接收
		fmt.Println(selectCase)
		//[{2 <chan int Value> <invalid Value>}]
		
		var selectCase1 = make([]reflect.SelectCase, 1)
		selectCase1[0].Dir = reflect.SelectSend //1
		selectCase1[0].Send = reflect.ValueOf(chs) //发送
		fmt.Println(selectCase1)
		//[{1 <invalid Value> <chan int Value>}]
		
		var selectCase2 = make([]reflect.SelectCase, 1)
		selectCase2[0].Dir = reflect.SelectDefault //1
		fmt.Println(selectCase2)
		//[{3 <invalid Value> <invalid Value>}]
    }

