# ring包

## 概述:

ring包实现了环形双向链表的功能。

## 函数列表：

- [type Ring](Ring.md)
  - [func New(n int) *Ring](New.md)
  - [func (r *Ring) Do(f func(interface{}))](Do.md)
  - [func (r *Ring) Len() int](Len.md)
  - [func (r *Ring) Link(s *Ring) *Ring](Link.md)
  - [func (r *Ring) Move(n int) *Ring](Move.md)
  - [func (r *Ring) Next() *Ring](Next.md)
  - [func (r *Ring) Prev() *Ring](Prev.md)
  - [func (r *Ring) Unlink(n int) *Ring](Unlink.md)
