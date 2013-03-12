# rand包函数列表

函数列表

- [func ExpFloat64() float64](ExpFloat64.md)
- [func Float32() float32](Float32.md)
- [func Float64() float64](Float64.md)
- [func Int() int](Int.md)
- [func Int31() int32](Int31.md)
- [func Int31n(n int32) int32](Int31n.md)
- [func Int63() int64](Int63.md)
- [func Int63n(n int64) int64](Int63n.md)
- [func Intn(n int) int](Intn.md)
- [func NormFloat64() float64](NormFloat64.md)
- func Perm(n int) []int]
- func Seed(seed int64)
- func Uint32() uint32
- type Rand
-    func New(src Source) *Rand
-    func (r *Rand) ExpFloat64() float64
-    func (r *Rand) Float32() float32
-    func (r *Rand) Float64() float64
-    func (r *Rand) Int() int
-    func (r *Rand) Int31() int32
-    func (r *Rand) Int31n(n int32) int32
-    func (r *Rand) Int63() int64
-    func (r *Rand) Int63n(n int64) int64
-    func (r *Rand) Intn(n int) int
-    func (r *Rand) NormFloat64() float64
-    func (r *Rand) Perm(n int) []int
-    func (r *Rand) Seed(seed int64)
-    func (r *Rand) Uint32() uint32
- type Source
-    func NewSource(seed int64) Source
- type Zipf
-    func NewZipf(r *Rand, s float64, v float64, imax uint64) *Zipf
-    func (z *Zipf) Uint64() uint64