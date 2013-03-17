# list包

## 概述:

list包实现了双向链表的功能。

遍历一个list的代码实例（其中`l`为*list对象）：

```go

	for e := l.Front(); e != nil; e = e.Next() {
		// do something with e.Value
	}

```

## 结构体：

- [type Element](Element.md)
  - [func (e *Element) Next() *Element](Next.md)
  - [func (e *Element) Prev() *Element](Prev.md)
- [type List](List.md)
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