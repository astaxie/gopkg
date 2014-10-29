# 项目发起的目的
由于目前golang的手册里面针对函数的例子太少了，很多时候不知道怎么使用，好多人都是看源代码才明白怎么用，这个给我们快速开发golang带来了障碍，所以我想发起这样一个项目，通过对pkg里面的针对每个函数写代码例子

# 格式要求

目录名表示包名，文件名是函数

例如strings目录表示strings这个包

strings目录下Contains.md表示strings.Contains这个函数的例子，代码例子参考我写的Contains.md 

样例请看：[strings](https://github.com/astaxie/gopkg/tree/master/strings)

代码必须在本机或者`http://play.golang.org`测试通过并且格式化

# 如何参与
如果你想参与这个项目，那么首先fork这个项目，然后修改[todo.md](https://github.com/astaxie/gopkg/blob/master/todo.md)文件，把你需要翻译的目录包填写在这个上面，然后pull给我，这样我们大家相互之间就不会做同一件事情了，不会导致两个人做了同一个包，这里面的包我每个目录已经建好了，但是需要注意：

>如果你填写了net包，那么不需要做下面的http这些子包的工作，这些也需要分包出去，每个人只做自己的包，不包括子包

欢迎你加群：386056972 和我们一起交流学习进步

# 贡献者

都在todo里面，谢谢这么多人的参与

