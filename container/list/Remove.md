# func (l *List) Remove(e *Element) interface{}

�����б�

- `e`������ɾ���Ľڵ㣬�ýڵ��������������`l`��

����ֵ��

- `interface{}`����ɾ���Ľڵ������

����˵����

ɾ��ָ���Ľڵ㣬����������ڵ������

����ʵ����

```go

	package main

	import (
		"fmt"
		"container/list"
	)

	func main(){
		l1 := list.New()
		l1.PushBack("a")
		l1.PushBack("b")
		fmt.Println(l1.Len()) // �����2
		
		l2 := list.New()
		l2.PushBack("c")
		elementFromL2 := l2.PushBack("d")
		
		l1.PushBackList(l2) // l2�����нڵ��list�ֶζ���l2����l2�����нڵ㶼�ӵ�l1��ĩβ��list�ֶα����l1
		fmt.Println(l1.Len()) // �����4
		
		elementFromL1 := l1.Back();
		fmt.Println(elementFromL1.Value) // �����d
		l1.Remove(elementFromL2) // elementFromL2������l2�ģ�������l1������l1�����нڵ㱻ɾ��
		fmt.Println(l1.Len()) // �����4
		fmt.Println(l1.Back().Value) // �����d
		l1.Remove(elementFromL1) // �ɹ���ɾ��
		fmt.Println(l1.Len()) // �����3
		fmt.Println(l1.Back().Value) // �����c
	}

```