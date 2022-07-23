package list

type SinglyLinked[T any] struct {
	Value T
	Next  *SinglyLinked[T]
}

type SinglyLinkedWithKey[K comparable, V any] struct {
	Key   K
	Value V
	Next  *SinglyLinkedWithKey[K, V]
}
