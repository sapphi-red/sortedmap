package sortedmap

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type NoLockSortedSet[K constraints.Ordered] struct {
	values []K
}

func NewNoLockSortedSet[K constraints.Ordered](capacity int) *NoLockSortedSet[K] {
	return &NoLockSortedSet[K]{
		values: make([]K, 0, capacity),
	}
}

func (s *NoLockSortedSet[K]) Size() int {
	return len(s.values)
}

func (s *NoLockSortedSet[K]) Capacity() int {
	return cap(s.values)
}

func (s *NoLockSortedSet[K]) ExtendCapacityTo(newCap int) {
	if s.Capacity() < newCap {
		s.values = append(make([]K, 0, newCap), s.values...)
	}
}

func (s *NoLockSortedSet[K]) Clear() {
	s.values = s.values[:0]
}

func (s *NoLockSortedSet[K]) Insert(value K) int {
	pos, exists := slices.BinarySearch(s.values, value)
	if exists {
		return -1
	}

	s.values = insertAt(s.values, pos, value)
	return pos
}

// TODO
// func (s *NewNoLockSortedSet[K]) InsertWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // inserted index
// }

func (s *NoLockSortedSet[K]) InsertWithAfterHint(value K, afterIndex int) int {
	partialValues := s.values[afterIndex:]
	pos, exists := slices.BinarySearch(partialValues, value)
	if exists {
		return -1
	}

	actualPos := afterIndex + pos
	s.values = insertAt(s.values, actualPos, value)
	return actualPos
}

func (s *NoLockSortedSet[K]) Delete(value K) int {
	pos, exists := slices.BinarySearch(s.values, value)
	if !exists {
		return -1
	}

	s.values = deleteAt(s.values, pos)
	return pos
}

// TODO
// func (s *NewNoLockSortedSet[K]) DeleteWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // deleted index
// }

func (s *NoLockSortedSet[K]) DeleteWithAfterHint(value K, afterIndex int) int {
	partialValues := s.values[afterIndex:]
	pos, exists := slices.BinarySearch(partialValues, value)
	if !exists {
		return -1
	}

	actualPos := afterIndex + pos
	s.values = deleteAt(s.values, actualPos)
	return actualPos
}

func (s *NoLockSortedSet[K]) InsertAll(values []K) {
	s.ExtendCapacityTo(s.Size() + len(values))

	for i := range values {
		s.Insert(values[i])
	}
}

func (s *NoLockSortedSet[K]) InsertAllOrdered(values []K) {
	s.ExtendCapacityTo(s.Size() + len(values))

	hint := 0
	for i := range values {
		hint = s.InsertWithAfterHint(values[i], hint)
	}
}

func (s *NoLockSortedSet[K]) DeleteAll(values []K) {
	s.ExtendCapacityTo(s.Size() + len(values))

	for i := range values {
		s.Delete(values[i])
	}
}

func (s *NoLockSortedSet[K]) DeleteAllOrdered(values []K) {
	s.ExtendCapacityTo(s.Size() + len(values))

	hint := 0
	for i := range values {
		hint = s.DeleteWithAfterHint(values[i], hint)
	}
}

func (s *NoLockSortedSet[K]) Contains(value K) bool {
	_, exists := slices.BinarySearch(s.values, value)
	return exists
}

func (s *NoLockSortedSet[K]) GetIndexOfGreater(value K) int {
	pos, exists := slices.BinarySearch(s.values, value)
	if exists {
		pos++ // does not include multiple same values
	}
	return pos
}
func (s *NoLockSortedSet[K]) GetIndexOfGreaterOrEqual(value K) int {
	pos, _ := slices.BinarySearch(s.values, value)
	return pos
}

func (s *NoLockSortedSet[K]) GetGreater(value K) []K {
	pos := s.GetIndexOfGreater(value)
	return s.values[pos:]
}
func (s *NoLockSortedSet[K]) GetGreaterOrEqual(value K) []K {
	pos := s.GetIndexOfGreaterOrEqual(value)
	return s.values[pos:]
}
func (s *NoLockSortedSet[K]) GetLess(value K) []K {
	pos := s.GetIndexOfGreaterOrEqual(value)
	return s.values[:pos]
}
func (s *NoLockSortedSet[K]) GetLessOrEqual(value K) []K {
	pos := s.GetIndexOfGreater(value)
	return s.values[:pos]
}

func (s *NoLockSortedSet[K]) GetByInclusiveRange(startValue K, endValue K) []K {
	startPos := s.GetIndexOfGreaterOrEqual(startValue)
	endPos := s.GetIndexOfGreater(endValue)
	return s.values[startPos:endPos]
}
