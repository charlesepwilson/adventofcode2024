package utils

import "go/types"

type Set[T comparable] struct {
	_map map[T]types.Nil
}

func (s Set[T]) Contains(t T) bool {
	_, ok := s._map[t]
	return ok
}
func (s Set[T]) Add(t T) {
	s._map[t] = types.Nil{}
}
func (s Set[T]) Remove(t T) {
	delete(s._map, t)
}
func (s Set[T]) Len() int {
	return len(s._map)
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{_map: make(map[T]types.Nil)}
}
