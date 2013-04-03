# reflect包函数列表

- [func Copy(dst, src Value) int](Copy.md)&nbsp;&nbsp;&nbsp;&nbsp;// 复制Slice或Array
- [func DeepEqual(a1, a2 interface{}) bool](DeepEqual.md)&nbsp;&nbsp;&nbsp;&nbsp;// 平等比较
- [type ChanDir](ChanDir.md)&nbsp;&nbsp;&nbsp;&nbsp;// 代表信道类型方向
	- func (d ChanDir) String() string&nbsp;&nbsp;&nbsp;&nbsp;// 以字符形式打印出来

			const (
				RecvDir ChanDir             = 1 << iota // <-chan 信道读取
				SendDir                                 // chan<- 信道写入
				BothDir = RecvDir | SendDir             // chan   信道读取与写入
			)

- [type Kind](Kind.md)&nbsp;&nbsp;&nbsp;&nbsp;// 一种数据类型
	- func (k Kind) String() string&nbsp;&nbsp;&nbsp;&nbsp;// 以字符形式打印出来

			const (
				Invalid Kind = iota	// 无效
				Bool				// 布尔
				Int					// 整数（有符号）
				Int8				// 整数8位（有符号）
				Int16				// 整数16位（有符号）
				Int32				// 整数32位（有符号）
				Int64				// 整数64位（有符号）
				Uint				// 整数（无符号）
				Uint8				// 整数8（无符号）
				Uint16				// 整数16（无符号）
				Uint32				// 整数（无符号）
				Uint64				// 整数（无符号）
				Uintptr				// 整数（指针,无符号）
				Float32				// 浮点数32位
				Float64				// 浮点数64位
				Complex64			// 复数64位
				Complex128			// 复数128位
				Array				// 数组
				Chan				// 信道
				Func				// 函数
				Interface			// 接口
				Map					// 地图
				Ptr					// 指针
				Slice				// 切片
				String				// 字符
				Struct				// 结构
				UnsafePointer		// 安全指针
			)

- [type Method](Method.md)&nbsp;&nbsp;&nbsp;&nbsp;// 方法

		type Method struct {
			Name	string	// 方法的名称
			PkgPath string	// 方法的路径（包的路径）
			Type  	Type	// 方法类型（reflect.Type）
			Func  	Value	// 方法的值（reflect.Value）
			Index 	int		// 指数列，“方法集”中的此方法排在第几。
		}

- type SliceHeader&nbsp;&nbsp;&nbsp;&nbsp;// 是运行时表示切片。它不能被安全地使用，或者可移植。

		type SliceHeader struct {
			Data	uintptr	// 指针
			Len		int		// 长度
			Cap		int		// 容量
		}

- type StringHeader&nbsp;&nbsp;&nbsp;&nbsp;// 是运行时表示切片。它不能被安全地使用，或者可移植。

		type StringHeader struct {
			Data	uintptr	// 指针
			Len		int		// 长度
		}

- [type StructField](StructField.md)&nbsp;&nbsp;&nbsp;&nbsp;// 在一个Struct结构内StructField描述了一个单一的字段。

		type StructField struct {
			Name		string		// 字段名称
			PkgPath 	string		// 结构路径名
			Type		Type      	// 字段类型
			Tag			StructTag 	// 字段标签字符
			Offset		uintptr   	// 结构内的偏移量，以字节为单位
			Index		[]int     	// 指数序列，“字段集”中的此字段排在第几。Type.FieldByIndex
			Anonymous	bool      	// 判断是否是匿名的字段。（没有类型）
		}

- type StructTag&nbsp;&nbsp;&nbsp;&nbsp;// 字符串标记，在结构字段内。
	- func (tag StructTag) Get(key string) string&nbsp;&nbsp;&nbsp;&nbsp;// 返回Key键标记字符串的值
- [type Type](Type.md)&nbsp;&nbsp;&nbsp;&nbsp;// 类型
	- [func PtrTo(t Type) Type](PtrTo.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回元素 t 的指针类型
	- [func TypeOf(i interface{}) Type](TypeOf.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回反射interface{}接口的类型
	- [NumMethod() int](Type.NumMethod.md)&nbsp;&nbsp;&nbsp;&nbsp;// 函数总数量，在struct结构中
	- [Method(int) Method](Type.Method.md)&nbsp;&nbsp;&nbsp;&nbsp;// 指定返回函数的 Method 类型，在struct结构中
	- [MethodByName(string) (Method, bool)](Type.MethodByName.md)&nbsp;&nbsp;&nbsp;&nbsp;// 使用“字符串”函数名称返回函数的 Method 类型，在struct结构中
	- [NumField() int](Type.NumField.md)&nbsp;&nbsp;&nbsp;&nbsp;// 字段总数量，在struct结构中
	- [Field(i int) StructField](Type.Field.md)&nbsp;&nbsp;&nbsp;&nbsp;// 指定返回字段的 StructField 类型，在struct结构中
	- [FieldByIndex(index []int) StructField](Type.FieldByIndex.md)&nbsp;&nbsp;&nbsp;&nbsp;// 指定返回“嵌套”字段的 StructField 类型，在struct结构中
	- [FieldByName(name string) (StructField, bool)](Type.FieldByName.md)&nbsp;&nbsp;&nbsp;&nbsp;// 使用“字符串”字段名称返回字段的 StructField 类型，在struct结构中
	- [FieldByNameFunc(match func(string) bool) (StructField, bool)](Type.FieldByNameFunc.md)&nbsp;&nbsp;&nbsp;&nbsp;// 传入字段“字符串”名称，并判断，func 返回 true，返回字段的 StructField 类型，在struct结构中
	- [NumIn() int](Type.NumIn.md)&nbsp;&nbsp;&nbsp;&nbsp;// 函数输入参数总数量
	- [In(i int) Type](Type.In.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回函数输入参数的第i个类型 Type
	- [NumOut() int](Type.NumOut.md)&nbsp;&nbsp;&nbsp;&nbsp;// 函数输出参数总数量
	- [Out(i int) Type](Type.Out.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回函数输出参数的第i个类型 Type
	- [Align() int](Type.Align.md)&nbsp;&nbsp;&nbsp;&nbsp;// 在分配在内存时的此类型的一个值（以字节为单位）的对齐。
	- [FieldAlign() int](Type.FieldAlign.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回字段对齐的值（以字节为单位）
	- [Name() string](Type.Name.md)&nbsp;&nbsp;&nbsp;&nbsp;// 变量名称或字段的名称
	- [PkgPath() string](Type.PkgPath.md)&nbsp;&nbsp;&nbsp;&nbsp;// 变量的（包）路径名
	- [Size() uintptr](Type.Size.md)&nbsp;&nbsp;&nbsp;&nbsp;// 值的数据大小（以字节为单位）
	- [String() string](Type.String.md)&nbsp;&nbsp;&nbsp;&nbsp;// （包）路径名称+类型名称
	- [Kind() Kind](Type.Kind.md)&nbsp;&nbsp;&nbsp;&nbsp;// 变量的类型
	- [Implements(u Type) bool](Type.Implements.md)&nbsp;&nbsp;&nbsp;&nbsp;// 判断是否存在与 u 相同的接口
	- [AssignableTo(u Type) bool](Type.AssignableTo.md)&nbsp;&nbsp;&nbsp;&nbsp;// 判断值是否可分配给 u
	- [Bits() int](Type.Bits.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回类型比特的大小
	- [ChanDir() ChanDir](Type.ChanDir.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回信道的方向
	- [IsVariadic() bool](Type.IsVariadic.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回函数的类型最后一个输入参数是否是“...”参数。
	- [Elem() Type](Type.Elem.md)&nbsp;&nbsp;&nbsp;&nbsp;// 指针指向内存地址
	- [Key() Type](Type.Key.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Map 键Key的类型
	- [Len() int](Type.Len.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Array 的长度
- [type Value](Value.md)&nbsp;&nbsp;&nbsp;&nbsp;// 值
	- [func Append(s Value, x ...Value) Value](Append.md)&nbsp;&nbsp;&nbsp;&nbsp;// 追加Slice
	- [func AppendSlice(s, t Value) Value](AppendSlice.md)&nbsp;&nbsp;&nbsp;&nbsp;// 批量追加Slice
	- [func Indirect(v Value) Value](Indirect.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回指针源内存地址
	- [func MakeChan(typ Type, buffer int) Value](MakeChan.md)&nbsp;&nbsp;&nbsp;&nbsp;// 初始化信道
	- [func MakeMap(typ Type) Value](MakeMap.md)&nbsp;&nbsp;&nbsp;&nbsp;// 初始化Map
	- [func MakeSlice(typ Type, len, cap int) Value](MakeSlice.md)&nbsp;&nbsp;&nbsp;&nbsp;// 初始化Slice
	- [func New(typ Type) Value](New.md)&nbsp;&nbsp;&nbsp;&nbsp;// 初始化并返回指针
	- [func NewAt(typ Type, p unsafe.Pointer) Value](NewAt.md)&nbsp;&nbsp;&nbsp;&nbsp;// 初始化 p 并返回指针，转向给 typ
	- [func Zero(typ Type) Value](Zero.md)&nbsp;&nbsp;&nbsp;&nbsp;// 初始化 typ 为零值
	- [func ValueOf(i interface{}) Value](ValueOf.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回反射interface{}接口的值
	- [func (v Value) Elem() Value](Value.Elem.md)&nbsp;&nbsp;&nbsp;&nbsp;// 指针指向内存地址
	- [func (v Value) Type() Type](Value.Type.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回类型 rflect.Type
	- [func (v Value) NumField() int](Value.NumField.md)&nbsp;&nbsp;&nbsp;&nbsp;// 字段总数量，在struct结构中
	- [func (v Value) Field(i int) Value](Value.Field.md)&nbsp;&nbsp;&nbsp;&nbsp;// 指定返回字段的 Value 类型，在struct结构中
	- [func (v Value) FieldByIndex(index []int) Value](Value.FieldByIndex.md)&nbsp;&nbsp;&nbsp;&nbsp;// 指定返回“嵌套”字段的 Value 类型，在struct结构中
	- [func (v Value) FieldByName(name string) Value](Value.FieldByName.md)&nbsp;&nbsp;&nbsp;&nbsp;// 使用“字符串”字段名称返回字段的 Value 类型，在struct结构中
	- [func (v Value) FieldByNameFunc(match func(string) bool) Value](Value.FieldByNameFunc.md)&nbsp;&nbsp;&nbsp;&nbsp;// 传入字段“字符串”名称，并判断，func 返回 true，返回字段的 Value 类型，在struct结构中
	- [func (v Value) NumMethod() int](Value.NumMethod.md)&nbsp;&nbsp;&nbsp;&nbsp;// 函数总数量，在struct结构中
	- [func (v Value) Method(i int) Value](Value.Method.md)&nbsp;&nbsp;&nbsp;&nbsp;// 指定返回函数的 Value 类型，在struct结构中
	- [func (v Value) MethodByName(name string) Value](Value.MethodByName.md)&nbsp;&nbsp;&nbsp;&nbsp;// 使用“字符串”函数名称返回函数的 Value 类型，在struct结构中
	- [func (v Value) Index(i int) Value](Value.Index.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回Array或Slice类型的第i个切片，在struct结构中
	- [func (v Value) Kind() Kind](Value.Kind.md)&nbsp;&nbsp;&nbsp;&nbsp;// 值的类型
	- [func (v Value) Call(in []Value) []Value](Value.Call.md)&nbsp;&nbsp;&nbsp;&nbsp;// 调用函数，in 切片装入参数，传入函数
	- [func (v Value) CallSlice(in []Value) []Value](Value.CallSlice.md)&nbsp;&nbsp;&nbsp;&nbsp;// 调用函数，in 切片装入参数，传入函数。用于可变参数函数
	- [func (v Value) IsValid() bool](Value.IsValid.md)&nbsp;&nbsp;&nbsp;&nbsp;// 判断值是否是零值
	- [func (v Value) IsNil() bool](Value.IsNil.md)&nbsp;&nbsp;&nbsp;&nbsp;// 判断值是否是 nil，限制支持 Chan，Func，Interface，Map，Ptr，或Slice
	- [func (v Value) CanInterface() bool](Value.CanInterface.md)&nbsp;&nbsp;&nbsp;&nbsp;// 判断值是否可以做为 interface{} 类型读出
	- [func (v Value) Interface() (i interface{})](Value.Interface.md)&nbsp;&nbsp;&nbsp;&nbsp;// 以接口类型读出数据
	- [func (v Value) InterfaceData() [2]uintptr](Value.InterfaceData.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回一对作为uintptr的接口值
	- [func (v Value) Slice(beg, end int) Value](Value.Slice.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回指定长度的切片
	- [func (v Value) CanAddr() bool](Value.CanAddr.md)&nbsp;&nbsp;&nbsp;&nbsp;// 判断是否是可以寻址
	- [func (v Value) Addr() Value](Value.Addr.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回指针值的地址
	- [func (v Value) UnsafeAddr() uintptr](Value.UnsafeAddr.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回安全指针指向v的数据
	- [func (v Value) CanSet() bool](Value.CanSet.md)&nbsp;&nbsp;&nbsp;&nbsp;// 判断是否可以写入值
	- [func (v Value) Set(x Value)](Value.Set.md)&nbsp;&nbsp;&nbsp;&nbsp;// 写入新值，支持所有类型
	- [func (v Value) Bool() bool](Value.Bool.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Bool 类型的值
	- [func (v Value) SetBool(x bool)](Value.SetBool.md)&nbsp;&nbsp;&nbsp;&nbsp;// 写入 Bool 类型的值
	- [func (v Value) Bytes() []byte](Value.Bytes.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Byte 类型的值
	- [func (v Value) SetBytes(x []byte)](Value.SetBytes.md)&nbsp;&nbsp;&nbsp;&nbsp;// 写入 Byte 类型的值
	- [func (v Value) Int() int64](Value.Int.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Int 类型的值
	- [func (v Value) OverflowInt(x int64) bool](Value.OverflowInt.md)&nbsp;&nbsp;&nbsp;&nbsp;// 判断 Int 类型的值承受范围
	- [func (v Value) SetInt(x int64)](Value.SetInt.md)&nbsp;&nbsp;&nbsp;&nbsp;// 写入 Int 类型的值
	- [func (v Value) Uint() uint64](Value.Uint.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Uint 类型的值
	- [func (v Value) OverflowUint(x uint64) bool](Value.OverflowUint.md)&nbsp;&nbsp;&nbsp;&nbsp;// 判断 Uint 类型的值承受范围
	- [func (v Value) SetUint(x uint64)](Value.SetUint.md)&nbsp;&nbsp;&nbsp;&nbsp;// 写入 Uint 类型的值
	- [func (v Value) Float() float64](Value.Float.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Float 类型的值
	- [func (v Value) OverflowFloat(x float64) bool](Value.OverflowFloat.md)&nbsp;&nbsp;&nbsp;&nbsp;// 判断 Float 类型的值承受范围
	- [func (v Value) SetFloat(x float64)](Value.SetFloat.md)&nbsp;&nbsp;&nbsp;&nbsp;// 写入 Float 类型的值
	- [func (v Value) Complex() complex128](Value.Complex.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Complex 类型的值
	- [func (v Value) OverflowComplex(x complex128) bool](Value.OverflowComplex.md)&nbsp;&nbsp;&nbsp;&nbsp;// 判断 Complex 类型的值承受范围
	- [func (v Value) SetComplex(x complex128)](Value.SetComplex.md)&nbsp;&nbsp;&nbsp;&nbsp;// 写入 Complex 类型的值
	- [func (v Value) Pointer() uintptr](Value.Pointer.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回指针（整数）
	- [func (v Value) SetPointer(x unsafe.Pointer)](Value.SetPointer.md)&nbsp;&nbsp;&nbsp;&nbsp;// 写入新的指针
	- [func (v Value) String() string](Value.String.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 String 类型的值
	- [func (v Value) SetString(x string)](Value.SetString.md)&nbsp;&nbsp;&nbsp;&nbsp;// 写入 String 类型的值
	- [func (v Value) MapKeys() []Value](Value.MapKeys.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Map 中的所有 Key 名称
	- [func (v Value) MapIndex(key Value) Value](Value.MapIndex.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Map 中 Key 的值
	- [func (v Value) SetMapIndex(key, val Value)](Value.SetMapIndex.md)&nbsp;&nbsp;&nbsp;&nbsp;// 设置 Mep 的 值
	- [func (v Value) Len() int](Value.Len.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Slice，Array，Chan，Map，String 长度
	- [func (v Value) Cap() int](Value.Cap.md)&nbsp;&nbsp;&nbsp;&nbsp;// 返回 Slice，Array，Chan 容量
	- [func (v Value) SetLen(n int)](Value.SetLen.md)&nbsp;&nbsp;&nbsp;&nbsp;// 改变 Slice 长度
	- [func (v Value) Recv() (x Value, ok bool)](Value.Recv.md)&nbsp;&nbsp;&nbsp;&nbsp;// 信道接收
	- [func (v Value) Send(x Value)](Value.Send.md)&nbsp;&nbsp;&nbsp;&nbsp;// 信道发送
	- [func (v Value) TryRecv() (x Value, ok bool)](Value.TryRecv.md)&nbsp;&nbsp;&nbsp;&nbsp;// 信道尝式接收
	- [func (v Value) TrySend(x Value) bool](Value.TrySend.md)&nbsp;&nbsp;&nbsp;&nbsp;// 信道尝式发送
	- [func (v Value) Close()](Value.Close.md)&nbsp;&nbsp;&nbsp;&nbsp;// 关闭信道
- type ValueError&nbsp;&nbsp;&nbsp;&nbsp;//调用方法不支持
	- func (e *ValueError) Error() string&nbsp;&nbsp;&nbsp;&nbsp;//返回错误内容
