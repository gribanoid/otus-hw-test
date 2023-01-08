package hw04lrucache

import (
	"sync"
)

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

func (l *list) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.len
}

func (l *list) Front() *ListItem {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.head
}

func (l *list) Back() *ListItem {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.mu.Lock()
	defer l.mu.Unlock()
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
	l.mu.Lock()
	defer l.mu.Unlock()
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
	l.mu.Lock()
	defer l.mu.Unlock()
	l.len--
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
}

func (l *list) MoveToFront(i *ListItem) {
	// undefined behavior in goroutines
	if i == l.head {
		return
	}
	l.Remove(i)
	l.PushFront(i.Value)

}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
	mu   sync.RWMutex
}

func NewList() List {
	return new(list)
}
