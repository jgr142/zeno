package set

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

func (s *Set[T]) Add(item T) {
	s.data[item] = struct{}{}
}

func (s *Set[T]) Remove(item T) {
	delete(s.data, item)
}

func (s *Set[T]) Contains(item T) bool {
	_, exists := s.data[item]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) Items() []T {
	items := make([]T, 0, len(s.data))
	for item := range s.data {
		items = append(items, item)
	}
	return items
}

func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}
