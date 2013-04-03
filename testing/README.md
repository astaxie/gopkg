# testing包

testing包为go的代码包提供自动化测试支持。它的目标是与“go test”命令协同使用，以自动执行form的任何函数。

    func TestXxx(*testing.T)

Xxx可以使任何字母和数字组合的字符串（但是第一个字符必须是a-z的）并且在测试程序中时唯一的。这些TestXxx程序应该在包中声明他们是用来测试的。

form的函数

   func BenchmarkXxx(*testing.B)

用来做基准测试，并且执行“go test -test.bench”命令后会启动这些测试。

一个简单的基准测试就像这样：

    func BenchmarkHello(b *testing.B) {
        for i := 0; i < b.N; i++ {
            fmt.Sprintf("hello")
        }
    }

基准测试包将会执行不同的 b.N 次直到基准测试函数可靠运行了足够长的时间。下面的输出：

    testing.BenchmarkHello    10000000    282 ns/op

意味着基准测试函数testing.BenchmarkHello()循环跑了10000000次，并且单次循环平均耗时 282 纳秒。

    func BenchmarkBigLen(b *testing.B) {
        b.StopTimer()
        big := NewBig()
        b.StartTimer()
        for i := 0; i < b.N; i++ {
            big.Len()
        }
    }

此包也可以运行和验证示例代码。示例测试函数可以包含结论注释语句，该语句以“Output:”开头，并当测试运行结束时标准输出会与这一注释语句相比较，下面是示例测试函数的例子：

    func ExampleHello() {
        fmt.Println("hello")
        // Output: hello
    }

    func ExampleSalutations() {
            fmt.Println("hello, and")
            fmt.Println("goodbye")
            // Output:
            // hello, and
            // goodbye
    }

没有输出注释语句的示例测试函数会被编译但不会被执行。

为函数 F、类型 T 以及类型 T 上的方法 M 声明示例测试函数的命名规则是这样的：

    func ExampleF() { ... }
    func ExampleT() { ... }
    func ExampleT_M() { ... }

为一个类型或函数或方法编写多个示例测试函数时可以直接在函数名称上加上不同的后缀。后缀必须以小写字母开头。

    func ExampleF_suffix() { ... }
    func ExampleT_suffix() { ... }
    func ExampleT_M_suffix() { ... }
    
只要测试文件中有示例测试函数，它就需要被明示是一个示例测试文件。这个示例测试文件可以包含其他函数、类型变量或常量声明，但不能有普通测试函数和基准测试函数。

# testing包函数列表

- [func Main(matchString func(pat, str string) (bool, error), tests []InternalTest, benchmarks []InternalBenchmark, examples []InternalExample)](Main.md)
- [func RunBenchmarks(matchString func(pat, str string) (bool, error), benchmarks []InternalBenchmark)](RunBenchmarks.md)
- [func RunExamples(matchString func(pat, str string) (bool, error), examples []InternalExample) (ok bool)](RunExamples.md)
- [func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)](RunTests.md)
- [func Short() bool](Short.md)
- [type B](B.md)
  - [func (c *B) Error(args ...interface{})](B_Error.md)
  - [func (c *B) Errorf(format string, args ...interface{})](B_Errorf.md)
  - [func (c *B) Fail()](B_Fail.md)
  - [func (c *B) FailNow()](B_FailNow.md)
  - [func (c *B) Failed() bool](B_Failed.md)
  - [func (c *B) Fatal(args ...interface{})](B_Fatal.md)
  - [func (c *B) Fatalf(format string, args ...interface{})](B_Fatalf.md)
  - [func (c *B) Log(args ...interface{})](B_Log.md)
  - [func (c *B) Logf(format string, args ...interface{})](B_Logf.md)
  - [func (b *B) ResetTimer()](B_ResetTimer.md)
  - [func (b *B) SetBytes(n int64)](B_SetBytes.md)
  - [func (b *B) StartTimer()](B_StartTimer.md)
  - [func (b *B) StopTimer()](B_StopTimer.md)
- [func Benchmark(f func(b *B)) BenchmarkResult](Benchmark.md)
- [type BenchmarkResult](BenchmarkResult.md)
  - [func (r BenchmarkResult) NsPerOp() int64](BenchmarkResult_NsPerOp.md)
  - [func (r BenchmarkResult) String() string](BenchmarkResult_String.md)
- [type InternalBenchmark](InternalBenchmark.md)
- [type InternalExample](InternalExample.md)
- [type InternalTest](InternalTest.md)
- [type T](T.md)
  - [func (c *T) Error(args ...interface{})](T_Error.md)
  - [func (c *T) Errorf(format string, args ...interface{})](T_Errorf.md)
  - [func (c *T) Fail()](T_Fail.md)
  - [func (c *T) FailNow()](T_FailNow.md)
  - [func (c *T) Failed() bool](T_Failed.md)
  - [func (c *T) Fatal(args ...interface{})](T_Fatal.md)
  - [func (c *T) Fatalf(format string, args ...interface{})](T_Fatalf.md)
  - [func (c *T) Log(args ...interface{})](T_Log.md)
  - [func (c *T) Logf(format string, args ...interface{})](T_Logf.md)
  - [func (t *T) Parallel()](T_Parallel.md)

