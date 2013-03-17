# type Element

## 结构体代码：

```go

	type Element struct {
		next, prev *Element

		list *List

		Value interface{}
	}

```

## 功能说明：

链表中的一个节点。

## 结构体字段：

- `next *Element`：指向链表中的下一个节点的指针，最后一个节点的下一个节点为`nil`
- `prev *Element`：指向链表中的上一个节点的指针，第一个节点的上一个节点为`nil`
- `list *List`：指向当前这个节点所属的链表的指针
- `Value interface{}`: 该节点的内容，可以是任何对象

## 函数链表：

- [func (e *Element) Next() *Element](Next.md)
- [func (e *Element) Prev() *Element](Prev.md)