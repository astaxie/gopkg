# func Copy(dst, src Value) int
参数列表

- dst Value 是目标切片Slice或数组Array
- src Value 是源切片Slice或数组Array

返回值：

- 返回 int 复制过去元素的数量

功能说明：

- Copy 复制src的内容复制到dst，直到dst已被填补满，或src已经耗尽。它返回复制的元素的数量。每个 dst 和 src 的 Kind（样）都必须切片(Slice)“或”数组(Array)，dst和src必须具有相同的元素类型。

代码实例：

    package main
    import (
        "fmt"
        "reflect"
    )
    
    type A struct {
      A0 []int
      A1 []int
    }
    
    func main(){
    	var a A
    	a.A0 = append(a.A0, []int{1,2,3,4,5,6,7}...)
    	a.A1 = append(a.A1, 9, 8, 7, 6)
    	var n = reflect.Copy(reflect.ValueOf(a.A0), reflect.ValueOf(a.A1))
    	fmt.Println(n, a)
    	//>>4 {[9 8 7 6 5 6 7] [9 8 7 6]}}
    }
