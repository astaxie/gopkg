## 接口 Interface

====
		type Interface interface {
			
			// Len 为集合内元素的总数
			Len() int
	   	 		
			// Less 返回索引为 i 的元素是否应排在索引为 j 的元素之前。
			Less(i, j int) bool
  	  			
			// Swap 交换索引为 i 和 j 的元素
			Swap(i, j int)
		}	

任何实现了 sort.Interface 的类型（一般为集合），均可使用该包中的方法进行排序。这些方法要求集合内列出元素的索引为整数。
		
代码案例（一）：
	
		package main
	
	import (
		"fmt"
		"sort"
	)
	
	type Grams int
	
	func (g Grams) String() string { return fmt.Sprintf("%dg", int(g)) }
	
	type Organ struct {
		Name   string
		Weight Grams
	}
	
	type Organs []*Organ
	
	func (s Organs) Len() int { return len(s) }
	
	func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
	
	// ByName 通过提供 Less 方法并使用嵌入式 Organs 值的 Len 和 Swap 方法实现了sort.Interface。
	
	type ByName struct{ Organs }
	
	func (s ByName) Less(i, j int) bool { return s.Organs[i].Name < s.Organs[j].Name }
	
	// ByWeight 通过提供 Less 方法并使用嵌入式 Organs 值的 Len 和 Swap 方法实现了sort.Interface。
	
	type ByWeight struct{ Organs }
	
	func (s ByWeight) Less(i, j int) bool { return s.Organs[i].Weight < s.Organs[j].Weight }
	
	func printOrgans(s []*Organ) {
		for _, o := range s {
			fmt.Printf("%-8s (%v)\n", o.Name, o.Weight)
		}
	}
	
	func main() {
		s := []*Organ{
			{"brain", 1340},
			{"heart", 290},
			{"liver", 1494},
			{"pancreas", 131},
		}
	
		sort.Sort(ByWeight{s})
		fmt.Println("Organs by weight:")
		printOrgans(s)
		// Output:
		// Organs by weight:
		// pancreas (131g)
		// heart    (290g)
		// brain    (1340g)
		// liver    (1494g)
		
		sort.Sort(ByName{s})
		fmt.Println("Organs by name:")
		printOrgans(s)
		// Output:
		// Organs by name:
		// brain    (1340g)
		// heart    (290g)
		// liver    (1494g)
		// pancreas (131g)	
	}
	
	
代码案例（二）：
	
	package main
		
	import (
		"fmt"
		"sort"
	)
	
	type Reverse struct {
		sort.Interface
	}
	
	// Less returns the opposite of the embedded implementation's Less method.
	
	func (r Reverse) Less(i, j int) bool {
		return r.Interface.Less(j, i)
	}
	
	func main() {
		s := []int{5, 2, 6, 3, 1, 4} // unsorted
		sort.Sort(Reverse{sort.IntSlice(s)})
		fmt.Println(s) // [6 5 4 3 2 1]
	}
												