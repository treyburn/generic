package queue

import (
	"github.com/treyburn/generic/list"
	"sync"
)

type TwoLockFIFO[T any] struct {
	headLock sync.Mutex
	tailLock sync.Mutex

	head *list.SinglyLinked[T]
	tail *list.SinglyLinked[T]
}

func NewFIFO[T any]() *TwoLockFIFO[T] {
	start := &list.SinglyLinked[T]{}
	return &TwoLockFIFO[T]{
		head: start,
		tail: start,
	}
}

func (tlq *TwoLockFIFO[T]) Enqueue(item T) {
	tlq.tailLock.Lock()
	defer tlq.tailLock.Unlock()

	newTail := &list.SinglyLinked[T]{
		Value: item,
		Next:  nil,
	}
	if tlq.tail != nil {
		tlq.tail.Next = newTail
	}
	tlq.tail = newTail
}

func (tlq *TwoLockFIFO[T]) Dequeue() T {
	tlq.headLock.Lock()
	defer tlq.headLock.Unlock()

	if tlq.head.Next == nil {
		return *new(T)
	}

	v := tlq.head.Next.Value
	tlq.head = tlq.head.Next

	return v
}
