#html/template包

##概述
html/template包实现了数据驱动的模板引擎,生成HTML输出并可防止代码注入. 虽然它提供了与text/template相同的接口,但如果你要输出的内容是HTML时应该使用html/template包而不是text/template包.	

这里着重介绍html/template包中的安全性功能. 有关如何编写自己的模板，请参阅text/template包。

##详细介绍
html/template包含了text/template,你可以使用template API来安全的解析和执行HTML模板.

	tmpl, err := template.New("name").Parse(...)
	// 这里省略了err的判断
	err = tmpl.Execute(out, data)
	
如果成功的话，tmpl已经过滤掉注入代码了,可以放心的使用.如果失败，err将返回一个ErrorCode.

在HTML模板中那些需要被转义或被编码的数据(值)会被当做纯文本进行处理,这样你就可以安全的把这些值嵌入到HTML中了.而且转义是语境关联的,JavaScript,CSS和URI中都会被转义.

template包使用的安全模型会假设模板的作者是可信的,而执行的数据和参数不是.更多细节请看下面的例子.

使用text/template包,结果不会被自动转义:	

	import "text/template"
	//这里省略一万字
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
	
输出的HTML没有经过转义: 	
`Hello, <script>alert('you have been pwned')</script>!`

而使用html/template包,输出会进行语境自动转义:

	import "html/template"
	//这里省略一万字
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
	
输出的HTML是经过转义的: 	
`Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!`

##语境转义
html/template包会转义的部分包含HTML, CSS, JavaScript, 和 URI	
它会为每一个管道添加消毒功能,下面是一个代码片段:	
`<a href="/search?q={{.}}">{{.}}</a>`	
在解析时,每一个`{{.}}`将会被覆盖,以增加必要的转义功能.在这种情况下，它将会变为:	
`<a href="/search?q={{. | urlquery}}">{{. | html}}</a>`	

##错误返回
详见文档中的ErrorCode部分

##更全面的了解
这个包其余的注释可以在第一次阅读时跳过,它包括了语境转义和错误异常返回的详细信息.大多数用户可能并不需要了解这些细节.

##语境转义详解

假设 `{{.}}` 为 `O'Reilly: How are <i>you</i>?`, 下面列出了转义之后`{{.}}`中的内容会变为什么

	语境                             {{.}} 之后
	
	{{.}}                            O'Reilly: How are &lt;i&gt;you&lt;/i&gt;?
	<a title='{{.}}'>                O&#39;Reilly: How are you?
	<a href="/{{.}}">                O&#39;Reilly: How are %3ci%3eyou%3c/i%3e?
	<a href="?q={{.}}">              O&#39;Reilly%3a%20How%20are%3ci%3e...%3f
	<a onx='f("{{.}}")'>             O\x27Reilly: How are \x3ci\x3eyou...?
	<a onx='f({{.}})'>               "O\x27Reilly: How are \x3ci\x3eyou...?"
	<a onx='pattern = /{{.}}/;'>     O\x27Reilly: How are \x3ci\x3eyou...\x3f
	
如果使用了一个不安全的内容,这个值将会被过滤掉:

	原内容                            {{.}} 之后
	
	<a href="{{.}}">                 #ZgotmplZ
	
例如"O'Reilly:" 是一个不安全的协议,如"http:".

如果 `{{.}}` 的内容是安全的, 例如`你好`, 他将不会被转义而是直接输出到模板中去

	语境                                 {{.}} 之后
	
	{{.}}                                你好
	<a title='{{.}}'>                    你好
	<a href='{{.}}'>                     你好
	<a href='/{{.}}'>                    你好
	<a href='?dir={{.}}'>                你好
	<a style="border-{{.}}: 4px">        你好
	<a style="align: {{.}}">             你好
	<a style="background: '{{.}}'>       你好
	<a style="background: url('{{.}}')>  你好
	<style>p.{{.}} {color:red}</style>   你好
	
非字符串的值是可以出现在JavaScript的语境中的. 例如 `{{.}}` 是:

	[]struct{A,B string}{ "foo", "bar" }
在模板转义过之后:

	<script>var pair = {{.}};</script>
模板的输出结果:

	<script>var pair = {"A": "foo", "B": "bar"};</script>
请参阅json包来了解如何在JavaScript语境中格式化的嵌入非字符串内容.

##字符串的类型
默认情况下,tmplate包假设所有的管道产生纯文本字符串.它增加了转义管道的阶段以确保在适当的情况下正确和安全地嵌入纯文本字符串.
当数据的值不是纯文本时,你可以确保它的标记不被转义.	
HTML, JS, URL等类型,并且内容是安全的,go语言可以使他们避免被转义.		
下面是模板的一个例子:	

	Hello, {{.}}!

可以被这样调用:	

	tmpl.Execute(out, HTML(`<b>World</b>`))

输出结果是:		

	Hello, <b>World</b>!

而不是:

	Hello, &lt;b&gt;World&lt;b&gt;!

如果`{{.}}`是一个普通的字符串将会被直接输出而不是被转义

##安全模型
template包使用的安全模型: 	
[http://js-quasis-libraries-and-repl.googlecode.com/svn/trunk/safetemplate.html#problem_definition](http://js-quasis-libraries-and-repl.googlecode.com/svn/trunk/safetemplate.html#problem_definition)	

template包使用的安全模型会假设模板的作者是可信的,而执行的数据和参数不是,并能在面对不受信任的数据时寻找并维持下面的这些特性：

- 结构保持性： 当使用安全的模板语言写一个HTML标签时,浏览器会解释并输出相应的部分,无论作为标记的值是否是不受信任的数据.同样,对于其他结构如属性界限,JS和CSS字符串的界限也是如此.

- 效果一致性： 注入模板的结果输出到一个页面,所有模板的使用者的运行应执行相同的结果.

- 最小惊讶性： 熟悉HTML,CSS和JavaScript的开发人员（或代码评审人员）,看到`{{.}}`时都应该知道自动语境转义发生了,并应该能正确推断出转义了什么.

##函数/类型列表
- [func HTMLEscape(w io.Writer, b []byte)](HTMLEscape.md)
- [func HTMLEscapeString(s string) string](HTMLEscapeString.md)
- [func HTMLEscaper(args ...interface{}) string](HTMLEscaper.md)
- func JSEscape(w io.Writer, b []byte)
- func JSEscapeString(s string) string
- func JSEscaper(args ...interface{}) string
- func URLQueryEscaper(args ...interface{}) string
- type CSS
- type Error
	- func (e *Error) Error() string
- type ErrorCode
- type FuncMap
- type HTML
- type HTMLAttr
- type JS
- type JSStr
- type Template
	- func Must(t *Template, err error) *Template
	- func New(name string) *Template
	- func ParseFiles(filenames ...string) (*Template, error)
	- func ParseGlob(pattern string) (*Template, error)
	- func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)
	- func (t *Template) Clone() (*Template, error)
	- func (t *Template) Delims(left, right string) *Template
	- func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
	- func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	- func (t *Template) Funcs(funcMap FuncMap) *Template
	- func (t *Template) Lookup(name string) *Template
	- func (t *Template) Name() string
	- func (t *Template) New(name string) *Template
	- func (t *Template) Parse(src string) (*Template, error)
	- func (t *Template) ParseFiles(filenames ...string) (*Template, error)
	- func (t *Template) ParseGlob(pattern string) (*Template, error)
	- func (t *Template) Templates() []*Template
- type URL
