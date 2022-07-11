package cache

import constant "github.com/treyburn/generic/const"

const ErrNotFound = constant.Err("key not found")

type doublyLinkedList[K comparable, V any] struct {
	key   K
	value V
	next  *doublyLinkedList[K, V]
	prev  *doublyLinkedList[K, V]
}

type LRUCache[K comparable, V any] struct {
	cap  int
	size int

	cache map[K]*doublyLinkedList[K, V]

	head *doublyLinkedList[K, V]
	tail *doublyLinkedList[K, V]
}

func NewLRU[K comparable, V any](capacity int) LRUCache[K, V] {
	c := make(map[K]*doublyLinkedList[K, V], capacity)

	head := doublyLinkedList[K, V]{}
	tail := doublyLinkedList[K, V]{}
	head.next = &tail
	tail.prev = &head

	return LRUCache[K, V]{cap: capacity, size: 0, cache: c, head: &head, tail: &tail}
}

func (lru *LRUCache[K, V]) addNode(node *doublyLinkedList[K, V]) {
	node.prev = lru.head
	node.next = lru.head.next

	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache[K, V]) removeNode(node *doublyLinkedList[K, V]) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

func (lru *LRUCache[K, V]) moveToHead(node *doublyLinkedList[K, V]) {
	lru.removeNode(node)
	lru.addNode(node)
}

func (lru *LRUCache[K, V]) popTail() *doublyLinkedList[K, V] {
	last := lru.tail.prev
	lru.removeNode(last)
	return last
}

func (lru *LRUCache[K, V]) Get(key K) (V, error) {
	v, ok := lru.cache[key]
	if !ok {
		return *new(V), ErrNotFound // this returns the zero value of type V - would it be better to return an error here?
	}
	lru.moveToHead(v)

	return v.value, nil
}

func (lru *LRUCache[K, V]) Put(key K, value V) {
	dll, ok := lru.cache[key]
	if !ok {
		dll = &doublyLinkedList[K, V]{
			key:   key,
			value: value,
			prev:  nil,
			next:  nil,
		}
		lru.addNode(dll)
		lru.size++
		lru.cache[key] = dll
		if lru.size > lru.cap {
			last := lru.popTail()
			delete(lru.cache, last.key)
			lru.size--
		}
	} else {
		dll.value = value
		lru.moveToHead(dll)
	}
}
