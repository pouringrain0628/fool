package set

type Set[T comparable] interface {
	Add(items ...T) Set[T]
	Contains(item T) bool
	ToList() []T
	Length() int
	Remove(items ...T) Set[T]
	Clear() Set[T]
	Union(other Set[T]) Set[T]
	Intersection(other Set[T]) Set[T]
	Difference(other Set[T]) Set[T]
}

type set[T comparable] struct {
	m map[T]struct{}
}

var EXIST = struct{}{}

func NewSet[T comparable](items ...T) Set[T] {
	s := &set[T]{}
	// 声明map类型的数据结构
	s.m = make(map[T]struct{})
	s.Add(items...)

	return s
}

func (s *set[T]) Add(items ...T) Set[T] {
	for _, item := range items {
		s.m[item] = EXIST
	}

	return s
}

func (s *set[T]) Contains(item T) bool {
	_, ok := s.m[item]
	return ok
}

func (s *set[T]) ToList() []T {
	var list []T

	for k := range s.m {
		list = append(list, k)
	}

	return list
}

func (s *set[T]) Length() int {
	return len(s.m)
}

func (s *set[T]) Remove(items ...T) Set[T] {
	for _, item := range items {
		delete(s.m, item)
	}

	return s
}

func (s *set[T]) Clear() Set[T] {
	s.m = make(map[T]struct{})
	return s
}

func (s *set[T]) Union(other Set[T]) Set[T] {
	unionSet := NewSet[T]()
	for item := range s.m {
		unionSet.Add(item)
	}

	for item := range other.(*set[T]).m {
		unionSet.Add(item)
	}

	return unionSet
}

func (s *set[T]) Intersection(other Set[T]) Set[T] {
	intersectionSet := NewSet[T]()

	for item := range s.m {
		if other.Contains(item) {
			intersectionSet.Add(item)
		}
	}

	return intersectionSet
}

func (s *set[T]) Difference(other Set[T]) Set[T] {
	differenceSet := NewSet[T]()

	for item := range s.m {
		if !other.Contains(item) {
			differenceSet.Add(item)
		}
	}

	return differenceSet
}
