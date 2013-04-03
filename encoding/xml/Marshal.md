##func Marshal(v interface{}) ([]byte, error)

参数列表:

- v 空接口，用于传入生产xml的数据

返回值:

- 生成的xml
- 是否有错误

功能说明:

当传入一个数组或者切片时，Marshal通过编组每一个元素来处理.当传入一个指针时，Marshal通过编组指针指向的值来处理，指针为空时，返回空.当传入一个接口值时，Marshal通过编组接口包含的值来处理接口，接口值为空，返回空.当传入其他数据时，返回一个或多个包含的数据。
xml元素名称的获取规则，按优先顺序排列:

    名称为XMLName字段的tag值 
    名称为XMLName并且类型为xml.Name的字段的值
    结构体名称

每一个struct包含的字段都会被作为输出的xml的元素，除了以下情况：

    名称为XMLName的字段将被忽略。
    tag为"-"的字段将被忽略。
    tag为"name,attr"的字段将变成xml元素中name命名的属性。
    tag为",chardata"的字段作为字符数据被写入，不是作为xml元素。
    tag为",innerxml"的字段逐字写入
    tag为",comment"的字段作为xml的注释，里面的数据必须不包括“--”.
    tag包含"omitempty"的字段，如果值为空值将会被忽略。空值包括：false,0,空接口，空指针，长度为零的数组、切片、map。
    值为外部结构体的字段，处理成非指针匿名结构字段

如果字段用了标签（tag）“a>b>c”，元素c将嵌套在他的父元素a和b中。父元素相同的字段，将封装在同一个元素内。
如果传入一个channel、function、map,Marshal将返回一个错误。

代码实例:
    
    package main
    
    import (
    	"encoding/xml"
    	"fmt"
    	"os"
    )
    
    type Address struct {
    	City, State string
    }
    type Person struct {
    	XMLName   xml.Name `xml:"person"`
    	Id        int      `xml:"id,attr"`
    	FirstName string   `xml:"name>first"`
    	LastName  string   `xml:"name>last"`
    	Age       int      `xml:"age"`
    	Height    float32  `xml:"height,omitempty"`
    	Married   bool
    	Address
    	Comment string `xml:",comment"`
    }
    
    func main() {
    	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
    	v.Comment = " Need more details. "
    	v.Address = Address{"Hanga Roa", "Easter Island"}
    
    	output, err := xml.Marshal(v)
    	if err != nil {
    		fmt.Printf("error: %v\n", err)
    	}
    
    	os.Stdout.Write(output)
    }
    

输出结果:

    <person id="13"><name><first>John</first><last>Doe</last></name><age>42</age><Married>false</Married><City>Hanga Roa</City><State>Easter Island</State><!-- Need more details. --></person>
    
