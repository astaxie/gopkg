# type Ring

## 结构体代码：

```go

	type Ring struct {
		next, prev *Ring
		Value      interface{}
	}

```

## 功能说明：

环形双向链表

## 结构体字段：

- `next *Ring`：指向链表中的下一个节点的指针
- `prev *Ring`：指向链表中的上一个节点的指针
- `Value interface{}`：该节点中存储的内容，可以是任何对象

## 函数链表：

- [func (r *Ring) Do(f func(interface{}))](Do.md)
- [func (r *Ring) Len() int](Len.md)
- [func (r *Ring) Link(s *Ring) *Ring](Link.md)
- [func (r *Ring) Move(n int) *Ring](Move.md)
- [func (r *Ring) Next() *Ring](Next.md)
- [func (r *Ring) Prev() *Ring](Prev.md)
- [func (r *Ring) Unlink(n int) *Ring](Unlink.md)
