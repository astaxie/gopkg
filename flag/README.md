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
- [func StringVar(p *string, name string, value string, usage string)](StringVar.md)
- [func Uint(name string, value uint, usage string) *uint](Uint.md)
- [func Uint64(name string, value uint64, usage string) *uint64](Uint64.md)
- [func Uint64Var(p *uint64, name string, value uint64, usage string)](Uint64Var.md)
- [func UintVar(p *uint, name string, value uint, usage string)](UintVar.md)
- [func Var(value Value, name string, usage string)](Var.md)
- [func Visit(fn func(*Flag))](Visit.md)
- [func VisitAll(fn func(*Flag))](VisitAll.md)

### type Flag
	
	type Flag struct {
	        Name     string // name as it appears on command line
	        Usage    string // help message
	        Value    Value  // value as set
	        DefValue string // default value (as text); for usage message
	}

- [func Lookup(name string) *Flag](Lookup.md)

### type FlagSet
    
    type FlagSet struct {
            Usage func()
    }

- [func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet](NewFlagSet.md)
- [func (f *FlagSet) Arg(i int) string](FlagSetArg.md)
- [func (f *FlagSet) Args() []string](FlagSetArgs.md)
- [func (f *FlagSet) Bool(name string, value bool, usage string) *bool](FlagSetBool.md)
- [func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)](FlagSetBoolVar.md)
- [func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration](FlagSetDuration.md)
- [func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)](FlagSetDurationVar.md)
- [func (f *FlagSet) Float64(name string, value float64, usage string) *float64](FlagSetFloat64.md)
- [func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)](FlagSetFloat64Var)
- [func (f *FlagSet) Init(name string, errorHandling ErrorHandling)](FlagSetInit.md)
- [func (f *FlagSet) Int(name string, value int, usage string) *int](FlagSetInt.md)
- [func (f *FlagSet) Int64(name string, value int64, usage string) *int64](FlagSetInt64.md)
- [func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)](FlagSetInt64Var.md)
- [func (f *FlagSet) IntVar(p *int, name string, value int, usage string)](FlagSetIntVar.md)
- [func (f *FlagSet) Lookup(name string) *Flag](FlagSetLookup.md)
- [func (f *FlagSet) NArg() int](FlagSetNArg.md)
- [func (f *FlagSet) NFlag() int](FlagSetNFlag.md)
- [func (f *FlagSet) Parse(arguments []string) error](FlagSetParse.md)
- [func (f *FlagSet) Parsed() bool](FlagSetParsed.md)
- [func (f *FlagSet) PrintDefaults()](FlagSetPrintDefaults.md)
- [func (f *FlagSet) Set(name, value string) error](FlagSetSet.md)
- [func (f *FlagSet) SetOutput(output io.Writer)](FlagSetSetOutput.md)
- [func (f *FlagSet) String(name string, value string, usage string) *string](FlagSetString.md)
- [func (f *FlagSet) StringVar(p *string, name string, value string, usage string)](FlagSetStringVar.md)
- [func (f *FlagSet) Uint(name string, value uint, usage string) *uint](FlagSetUint.md)
- [func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64](FlagSetUint64.md)
- [func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)](FlagSetUint64Var.md)
- [func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)](FlagSetUintVar.md)
- [func (f *FlagSet) Var(value Value, name string, usage string)](FlagSetVar.md)
- [func (f *FlagSet) Visit(fn func(*Flag))](FlagSetVisit.md)
- [func (f *FlagSet) VisitAll(fn func(*Flag))](FlagSetVisitAll.md)

### type Getter
    
    type Getter interface {
            Value
            Get() interface{}
    }

### type Value
    
    type Value interface {
            String() string
            Set(string) error
    }
