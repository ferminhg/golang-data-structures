package linked_lists

type Node[T any] struct {
	value T
	next  *Node[T]
}

type SinglyLinkedList[T any] struct {
	head *Node[T]
	size int
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: nil,
		size: 0,
	}
}

func (s *SinglyLinkedList[T]) Add(value T) {
	node := Node[T]{value: value}

	switch s.head {
	case nil:
		s.head = &node
	default:
		current := s.head
		for current.next != nil {
			current = current.next
		}

		current.next = &node
	}
	s.size++
}

func (s *SinglyLinkedList[T]) Length() int {
	return s.size
}

func (s *SinglyLinkedList[T]) Last() Node[T] {
	if s.head == nil {
		return Node[T]{}
	}

	current := s.head
	for current.next != nil {
		current = current.next
	}
	return *current
}

func (s *SinglyLinkedList[T]) First() Node[T] {
	return *s.head
}

func (s *SinglyLinkedList[T]) Get(index int) Node[T] {
	if index < 0 || index >= s.size {
		return Node[T]{}
	}

	current := s.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return *current
}
