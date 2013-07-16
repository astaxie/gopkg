# expvar包

expvar包提供了一组标准接口，将服务器内部的公共变量，通过HTTP协议，及json格式，暴露出来，可访问 /debug/vars 获取。

默认提供两个变量：

    cmdline os.Args
    memstats  runtime.Memstats

“公共变量”即Var，是一个实现了String()函数的接口，定义如下：

    type Var interface {
        String() string
    }

实际类型的Var包括：Int，Float，String和Map，每个具体的类型都包括几个函数：

- New*()      // 新建一个变量
- Set(*)      // 设置这个变量
- Add(*)      // 在原有变量上加上另一个变量
-	String()    // 实现Var接口

除此之外，Map还有几个特有的函数：

-	Init()                  // 初始化Map
-	Get(key string)         // 根据key获取value
-	Do(f func(KeyValue))    // 对Map中的每对key/value执行函数f

所有对Var的设置和修改都是原子操作，所有操作的实例在testExpvar.go, 运行后可访问 127.0.0.1/debug/vars, 一般会得到如vars.json的结果。
		
