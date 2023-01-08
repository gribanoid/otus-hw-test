package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.RWMutex
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.RLock()
	v, ok := l.items[key]
	l.mu.RUnlock()
	if ok {
		v.Value = value
		l.queue.MoveToFront(v)
		return true
	}
	listItem := l.queue.PushFront(value)
	l.mu.Lock()
	l.items[key] = listItem
	l.mu.Unlock()
	if l.queue.Len() > l.capacity {
		l.queue.Remove(l.queue.Back())
	}
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.RLock()
	v, ok := l.items[key]
	l.mu.RUnlock()
	if ok {
		l.queue.MoveToFront(v)
		return l.queue.Front().Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.items = map[Key]*ListItem{}
	l.queue = nil
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
