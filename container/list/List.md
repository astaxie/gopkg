# type List

## �ṹ����룺

```go

	type List struct {
		front, back *Element
		len         int
	}

```

## �ṹ���ֶΣ�

- `front *Element`�������еĵ�һ���ڵ��ָ��
- `back *Element`�����������һ���ڵ��ָ��
- `len int`�������нڵ�ĸ���

## ��������

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