package gs

type Stack[T any] struct {
	head *stackNode[T]
	len  int
}

type stackNode[T any] struct {
	value T
	next  *stackNode[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{head: nil, len: 0}
}

func (s *Stack[T]) Empty() bool {
	return s.len == 0
}

func (s *Stack[T]) Len() int {
	return s.len
}

func (s *Stack[T]) Push(value T) {
	s.len++
	s.head = &stackNode[T]{value: value, next: s.head}
}

func (s *Stack[T]) Pop() T {
	if s.len == 0 {
		panic("stack is empty")
	}
	s.len--
	value := s.head.value
	s.head = s.head.next

	return value
}
