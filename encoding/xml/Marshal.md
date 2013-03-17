##func Marshal(v interface{}) ([]byte, error)

参数列表:

- w 表示输出接口
- s 表示待转换的数据

返回值:

- 无

功能说明:

当传入一个数组或者切片时，Marshal通过编组每一个元素来处理.当传入一个指针时，Marshal通过编组指针指向的值来处理，指针为空时，返回空.当传入一个接口值时，Marshal通过编组接口包含的值来处理接口，接口值为空，返回空.当传入其他数据时，返回一个或多个包含的数据。
xml元素名称的获取规则，按优先顺序排列:

    名称为XMLName字段的tag 
    名称为XMLName并且类型为xml.Name的字段的值
    &   &amp;
    <   &lt;
    >   &gt;

每一个struct包含的字段都会被作为输出的xml的元素，除了以下情况：

    名称为XMLName的字段将被忽略。
    tag为"-"的字段将被忽略。
    tag为"name,attr"的字段将变成xml元素中name命名的属性。
    tag为",chardata"的字段作为字符数据被写入，不是作为xml元素。
    tag为",innerxml"的字段逐字写入

代码实例:

    package main
    
    import (
    	"encoding/xml"
    	"os"
    )
    
    func main() {
    	str := "\"'&<>"
    	xml.Escape(os.Stdout, []byte(str))
    }

输出结果:

    &#34;&#39;&amp;&lt;&gt;
