package cache

import (
	"github.com/treyburn/generic/constant"
	"github.com/treyburn/generic/list"
	"sync"
)

const ErrNotFound = constant.Err("key not found")

type LRUCache[K comparable, V any] struct {
	cap  int
	size int

	mu sync.Mutex

	cache map[K]*list.DoublyLinkedWithKey[K, V]

	head *list.DoublyLinkedWithKey[K, V]
	tail *list.DoublyLinkedWithKey[K, V]
}

func NewLRU[K comparable, V any](capacity int) *LRUCache[K, V] {
	c := make(map[K]*list.DoublyLinkedWithKey[K, V], capacity)

	head := list.DoublyLinkedWithKey[K, V]{}
	tail := list.DoublyLinkedWithKey[K, V]{}
	head.Next = &tail
	tail.Prev = &head

	return &LRUCache[K, V]{cap: capacity, size: 0, cache: c, head: &head, tail: &tail}
}

func (lru *LRUCache[K, V]) addNode(node *list.DoublyLinkedWithKey[K, V]) {
	node.Prev = lru.head
	node.Next = lru.head.Next

	lru.head.Next.Prev = node
	lru.head.Next = node
}

func (lru *LRUCache[K, V]) removeNode(node *list.DoublyLinkedWithKey[K, V]) {
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev
}

func (lru *LRUCache[K, V]) moveToHead(node *list.DoublyLinkedWithKey[K, V]) {
	lru.removeNode(node)
	lru.addNode(node)
}

func (lru *LRUCache[K, V]) popTail() *list.DoublyLinkedWithKey[K, V] {
	last := lru.tail.Prev
	lru.removeNode(last)
	return last
}

func (lru *LRUCache[K, V]) Get(key K) (V, error) {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	v, ok := lru.cache[key]
	if !ok {
		return *new(V), ErrNotFound
	}
	lru.moveToHead(v)

	return v.Value, nil
}

func (lru *LRUCache[K, V]) Put(key K, value V) {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	dll, ok := lru.cache[key]
	if !ok {
		dll = &list.DoublyLinkedWithKey[K, V]{
			Key:   key,
			Value: value,
			Prev:  nil,
			Next:  nil,
		}
		lru.addNode(dll)
		lru.size++
		lru.cache[key] = dll
		if lru.size > lru.cap {
			last := lru.popTail()
			delete(lru.cache, last.Key)
			lru.size--
		}
	} else {
		dll.Value = value
		lru.moveToHead(dll)
	}
}
