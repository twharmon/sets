package sets

type Ordered interface {
	int | int32 | int16 | int8 | int64 | uint | uint32 | uint16 | uint8 | uint64 | float32 | float64 | string
}

// Set holds a set of values.
type Set[T Ordered] struct {
	m map[T]struct{}
}

// New creates a new Set with the given initial values.
func New[T Ordered](v ...T) *Set[T] {
	s := Set[T]{
		m: make(map[T]struct{}),
	}
	for i := range v {
		s.m[v[i]] = struct{}{}
	}
	return &s
}

// Add adds a value to the Set.
func (s *Set[T]) Add(v ...T) {
	for i := range v {
		s.m[v[i]] = struct{}{}
	}
}

// Remove removes a value to the Set.
func (s *Set[T]) Remove(v ...T) {
	for i := range v {
		delete(s.m, v[i])
	}
}

// Contains returns true if the Set contains the given value and
// false otherwise.
func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

// Clear removes all values from the Set.
func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

// Len returns the number of values in the Set.
func (s *Set[T]) Len() int {
	return len(s.m)
}

// Slice returns a slice of all the values in the Set.
func (s *Set[T]) Slice() []T {
	output := make([]T, 0, len(s.m))
	for i := range s.m {
		output = append(output, i)
	}
	return output
}

// Merge adds all values from the given Sets to s.
func (s *Set[T]) Merge(v ...*Set[T]) {
	for i := range v {
		for j := range v[i].m {
			s.m[j] = struct{}{}
		}
	}
}

// Union returns a new Set with all the values found in any of the
// given Sets.
func Union[T Ordered](v ...*Set[T]) *Set[T] {
	s := New[T]()
	for i := range v {
		s.Merge(v[i])
	}
	return s
}

// Intersection returns a new Set with the values found in all of the
// given Sets.
func Intersection[T Ordered](v ...*Set[T]) *Set[T] {
	s := New[T]()
	if len(v) == 0 {
		return s
	}
	hash := make(map[T]int)
	for i := range v {
		for j := range v[i].m {
			if hash[j] == i {
				hash[j]++
			}
		}
	}
	for k := range hash {
		if hash[k] == len(v) {
			s.Add(k)
		}
	}
	return s
}
