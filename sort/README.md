# sort包函数列表

- [Float64s(a []float64)](Float64s.md)

- [Float64sAreSorted(a []float64) bool](Float64sAreSorted.md)

- [Ints(a []int)](Ints.md)

- [IntsAreSorted(a []int) bool](IntsAreSorted.md)

- [IsSorted(data Interface) bool](IsSorted.md)

- [Search(n int, f func(int) bool) int](Search.md)

- [SearchFloat64s(a []float64, x float64) int](SearchFloat64s.md)

- [SearchInts(a []int, x int) int](SearchInts.md)

- [SearchStrings(a []string, x string) int](SearchStrings.md)

- [Sort(data Interface)](Sort.md)

- [Strings(a []string)](Strings.md)

- [StringsAreSorted(a []string) bool](StringsAreSorted.md)
	
# 结构

- [type Interface](Interface.md)
	
- [type Float64Slice](Float64Slice.md)	
 - Len() int
 - Less(i, j int) bool
 - Search(x float64) int
 - Sort()
 - Swap(i, j int)
		
- [type IntSlice](IntSlice.md)
 - Len() int
 - Less(i, j int) bool
 - Search(x int) int
 - Sort()
 - Swap(i, j int)
			
- [type StringSlice](StringSlice.md)
 - Len() int
 - Less(i, j int) bool
 - Search(x string) int
 - Sort()
 - Swap(i, j int)
