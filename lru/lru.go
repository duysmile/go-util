package lru

import (
	"container/list"
	"sync"
)

type Key interface{}

type LRU struct {
	maxEntries int
	ll         *list.List
	cache      map[interface{}]*list.Element
	locker     sync.Mutex
}

type entry struct {
	key   Key
	value interface{}
}

func NewLRU(maxEntries int) *LRU {
	return &LRU{
		maxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

func (l *LRU) Add(key Key, value interface{}) {
	l.locker.Lock()
	defer l.locker.Unlock()

	if l.cache == nil {
		l.cache = make(map[interface{}]*list.Element)
		l.ll = list.New()
	}

	if e, ok := l.cache[key]; ok {
		e.Value = value
		l.ll.MoveToFront(e)
	} else {
		el := l.ll.PushFront(&entry{key, value})
		l.cache[key] = el
	}

	if l.maxEntries > 0 && l.ll.Len() > l.maxEntries {
		l.RemoveOldest()
	}
}

func (l *LRU) Get(key Key) (value interface{}, ok bool) {
	if l.cache == nil {
		return
	}
	if el, hit := l.cache[key]; hit {
		if entry, ok := el.Value.(*entry); ok {
			value = entry.value
			l.ll.MoveToFront(el)
			return value, ok
		}
	}

	return
}

func (l *LRU) Remove(key Key) {
	l.locker.Lock()
	defer l.locker.Unlock()

	if l.cache == nil {
		return
	}

	if el, ok := l.cache[key]; ok {
		l.removeElement(el)
	}
}

func (l *LRU) RemoveOldest() {
	if l.cache == nil {
		return
	}
	el := l.ll.Back()
	if el != nil {
		l.removeElement(el)
	}
}

func (l *LRU) removeElement(el *list.Element) {
	if e, ok := el.Value.(*entry); ok {
		delete(l.cache, e.key)
	}
	l.ll.Remove(el)
}
