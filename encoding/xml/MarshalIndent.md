##func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

参数列表:

- v 空接口，用于传入生产xml的数据
- prefix 字符串，xml根元素缩进
- indent 字符串，xml元素缩进

返回值:

- 生成的xml
- 是否有错误

功能说明:

生产原理同Marshal，根元素缩进prefix，其他元素缩进indent。

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
    
    	output, err := xml.MarshalIndent(v, "  ", "    ")
    	if err != nil {
    		fmt.Printf("error: %v\n", err)
    	}
    
    	os.Stdout.Write(output)
    }

输出结果:

      <person id="13">
          <name>
              <first>John</first>
              <last>Doe</last>
          </name>
          <age>42</age>
          <Married>false</Married>
          <City>Hanga Roa</City>
          <State>Easter Island</State>
          <!-- Need more details. -->
      </person>
 
