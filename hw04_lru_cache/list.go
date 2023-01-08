package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.len++
	front := &ListItem{Value: v}
	if l.len == 1 {
		l.head, l.tail = front, front
		return front
	}
	l.head.Prev, front.Next = front, l.head
	l.head = front
	return front
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.len++
	back := &ListItem{Value: v}
	if l.len == 1 {
		l.head, l.tail = back, back
		return back
	}
	l.tail.Next, back.Prev = back, l.tail
	l.tail = back
	return back
}

func (l *list) Remove(i *ListItem) {
	l.len--
	l.destroyElem(i)
}

func (l *list) MoveToFront(i *ListItem) {
	if i != l.head {
		l.destroyElem(i)
		front := &ListItem{Value: i.Value}
		if l.len == 1 {
			l.head, l.tail = front, front
		}
		l.head.Prev, front.Next = front, l.head
		l.head = front
	}
}

func (l *list) destroyElem(i *ListItem) {
	if l.len == 0 {
		l.head, l.tail = nil, nil
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head, i.Next.Prev = i.Next, nil
		return
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail, i.Prev.Next = i.Prev, nil
		return
	}
}
