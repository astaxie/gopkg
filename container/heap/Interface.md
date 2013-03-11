# type Interface

接口代码：

	type Interface interface {
	    sort.Interface
	    Push(x interface{})
	    Pop() interface{}
	}

接口方法：

- sort.Interface 

	参见sort.Interface接口

- Push(x interface{}) 

	参数列表：

	x 将存到堆中的元素

	功能说明：

	把x存放到索引号为Len()的位置上，比如，一个列表中元素有7个，索引号从0开始，那么x将被存放到索引号为7的位置上，即最末尾。

- Pop() interface{}
	
	返回值：

	被移除的索引号为Len()-1的元素

	功能说明：
	
	把索引号为Len()-1的元素移除并返回。比如，如果这个堆是一个数组，那么才操作就是把数组的最后一个元素移除并返回。

