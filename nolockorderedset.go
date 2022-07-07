package orderedmap

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type NoLockOrderedSet[K constraints.Ordered] struct {
	values []K
}

func NewNoLockOrderedSet[K constraints.Ordered](capacity int) *NoLockOrderedSet[K] {
	return &NoLockOrderedSet[K]{
		values: make([]K, 0, capacity),
	}
}

func (s *NoLockOrderedSet[K]) Size() int {
	return len(s.values)
}

func (s *NoLockOrderedSet[K]) Capacity() int {
	return cap(s.values)
}

func (s *NoLockOrderedSet[K]) ExtendCapacityTo(newCap int) {
	if s.Capacity() < newCap {
		s.values = append(make([]K, 0, newCap), s.values...)
	}
}

func (s *NoLockOrderedSet[K]) Clear() {
	s.values = s.values[:0]
}

func (s *NoLockOrderedSet[K]) Insert(value K) int {
	pos, exists := slices.BinarySearch(s.values, value)
	if exists {
		return -1
	}

	s.values = append(s.values, value)
	copy(s.values[pos+1:], s.values[pos:])
	s.values[pos] = value
	return pos
}

// TODO
// func (s *NewNoLockOrderedSet[K]) InsertWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // inserted index
// }

// TODO
// func (s *NewNoLockOrderedSet[K]) InsertWithAfterHint(value K, afterIndex int) int {
// 	return 0 // inserted index
// }

func (s *NoLockOrderedSet[K]) Delete(value K) int {
	pos, exists := slices.BinarySearch(s.values, value)
	if !exists {
		return -1
	}

	s.values = append(s.values[:pos], s.values[pos+1:]...)
	return pos
}

// TODO
// func (s *NewNoLockOrderedSet[K]) DeleteWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // deleted index
// }

// TODO
// func (s *NewNoLockOrderedSet[K]) DeleteWithAfterHint(value K, afterIndex int) int {
// 	return 0 // deleted index
// }

func (s *NoLockOrderedSet[K]) InsertAll(values []K) {
	s.ExtendCapacityTo(s.Size() + len(values))

	for i := range values {
		s.Insert(values[i])
	}
}

// TODO
// func (s *NewNoLockOrderedSet[K]) InsertAllOrdered(values []K) {
// }

func (s *NoLockOrderedSet[K]) Contains(value K) bool {
	_, exists := slices.BinarySearch(s.values, value)
	return exists
}

func (s *NoLockOrderedSet[K]) GetIndexOfGreater(value K) int {
	pos, exists := slices.BinarySearch(s.values, value)
	if exists {
		pos++ // does not include multiple same values
	}
	return pos
}
func (s *NoLockOrderedSet[K]) GetIndexOfGreaterOrEqual(value K) int {
	pos, _ := slices.BinarySearch(s.values, value)
	return pos
}

func (s *NoLockOrderedSet[K]) GetGreater(value K) []K {
	pos := s.GetIndexOfGreater(value)
	return s.values[pos:]
}
func (s *NoLockOrderedSet[K]) GetGreaterOrEqual(value K) []K {
	pos := s.GetIndexOfGreaterOrEqual(value)
	return s.values[pos:]
}
func (s *NoLockOrderedSet[K]) GetLess(value K) []K {
	pos := s.GetIndexOfGreaterOrEqual(value)
	return s.values[:pos]
}
func (s *NoLockOrderedSet[K]) GetLessOrEqual(value K) []K {
	pos := s.GetIndexOfGreater(value)
	return s.values[:pos]
}

func (s *NoLockOrderedSet[K]) GetByInclusiveRange(startValue K, endValue K) []K {
	startPos := s.GetIndexOfGreaterOrEqual(startValue)
	endPos := s.GetIndexOfGreater(endValue)
	return s.values[startPos:endPos]
}
