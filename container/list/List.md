# type List

## 结构体代码：

```go

	type List struct {
		front, back *Element
		len         int
	}

```

## 结构体字段：

- `front *Element`：链表中的第一个节点的指针
- `back *Element`：链表中最后一个节点的指针
- `len int`：链表中节点的个数

## 函数链表：

- [func New() *List](New.md)
- [func (l *List) Back() *Element](Back.md)
- [func (l *List) Front() *Element](Front.md)
- [func (l *List) Init() *List](Init.md)
- [func (l *List) InsertAfter(value interface{}, mark *Element) *Element](InsertAfter.md)
- [func (l *List) InsertBefore(value interface{}, mark *Element) *Element](InsertBefore.md)
- [func (l *List) Len() int](Len.md)
- [func (l *List) MoveToBack(e *Element)](MoveToBack.md)
- [func (l *List) MoveToFront(e *Element)](MoveToFront.md)
- [func (l *List) PushBack(value interface{}) *Element](PushBack.md)
- [func (l *List) PushBackList(ol *List)](PushBackList.md)
- [func (l *List) PushFront(value interface{}) *Element](PushFront.md)
- [func (l *List) PushFrontList(ol *List)](PushFrontList.md)
- [func (l *List) Remove(e *Element) interface{}](Remove.md)