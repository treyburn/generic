package list

type DoublyLinked[T any] struct {
	Value T
	Next  *DoublyLinked[T]
	Prev  *DoublyLinked[T]
}

type DoublyLinkedWithKey[K comparable, V any] struct {
	Key   K
	Value V
	Next  *DoublyLinkedWithKey[K, V]
	Prev  *DoublyLinkedWithKey[K, V]
}
