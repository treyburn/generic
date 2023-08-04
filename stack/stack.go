package stack

type Stack[T any] struct {
	store []T
}

func (s *Stack[T]) Push(item T) {
	s.store = append(s.store, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.store) == 0 {
		return *new(T)
	}

	item := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return item
}

func (s *Stack[T]) Peek() T {
	if len(s.store) == 0 {
		return *new(T)
	}

	return s.store[len(s.store)-1]
}

func (s *Stack[T]) ToArray() []T {
	return s.store
}

func (s *Stack[T]) Len() int {
	return len(s.store)
}

func New[T any]() *Stack[T] {
	s := make([]T, 0)
	return &Stack[T]{store: s}
}
