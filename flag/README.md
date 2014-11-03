# flag包使用详解

##概述
flag包提供了一系列解析命令行参数的功能接口

###Variables
    
    var (
      CommandLine = NewFlagSet(os.Args[0], ExitOnError)
      ErrHelp     = errors.New("flag: help requested")
      Usage       = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        PrintDefaults()
      }
    )


##包函数列表
- [func Arg(i int) string](Arg.md)
- [func Args() []string](Args.md)
- [func Bool(name string, value bool, usage string) *bool](Bool.md)
- [func BoolVar(p *bool, name string, value bool, usage string)](BoolVar.md)
- [func Duration(name string, value time.Duration, usage string) *time.Duration](Duration.md)
- [func DurationVar(p *time.Duration, name string, value time.Duration, usage string)](DurationVar.md)
- [func Float64(name string, value float64, usage string) *float64](Float64.md)
- [func Float64Var(p *float64, name string, value float64, usage string)](Float64Var.md)
- [func Int(name string, value int, usage string) *int](Int.md)
- [func Int64(name string, value int64, usage string) *int64](Int64.md)
- [func Int64Var(p *int64, name string, value int64, usage string)](Int64Var.md)
- [func IntVar(p *int, name string, value int, usage string)](IntVar.md)
- [func NArg() int](NArg.md)
- [func NFlag() int](NFlag.md)
- [func Parse()](Parse.md)
- [func Parsed() bool](Parsed.md)
- [func PrintDefaults()](PrintDefaults.md)
- [func Set(name, value string) error](Set.md)
- [func String(name string, value string, usage string) *string](String.md)
- [func StringVar(p *string, name string, value string, usage string)](StringVar)
- [func Uint(name string, value uint, usage string) *uint](Uint.md)
- [func Uint64(name string, value uint64, usage string) *uint64](Uint64.md)
- [func Uint64Var(p *uint64, name string, value uint64, usage string)](Uint64Var.md)
- [func UintVar(p *uint, name string, value uint, usage string)](UintVar.md)
- [func Var(value Value, name string, usage string)](Var.md)
- [func Visit(fn func(*Flag))](Visit.md)
- [func VisitAll(fn func(*Flag))](VisitAll.md)