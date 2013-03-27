# func MultiReader(readers ...Reader) Reader

参数：
- readers 不定参数，为若干读取器

返回值：MultiReader

MultiReader说明：
- 封装多个Reader，可实现多个Reader连续读取，就想在读取一个Reader一样

功能说明：
- 获得一个可以对参数中的多个Reader进行连续读取的Reader

示例：
  package main
	
	import (
		"io"
		"fmt"
		"os"
		"reflect"
	)
	
	func main() {
		reader1, _ := os.Open("src1.txt")
		reader2, _ := os.Open("src2.txt")
		multiReader := io.MultiReader(reader1, reader2)
		fmt.Println(reflect.TypeOf(multiReader))
	}
