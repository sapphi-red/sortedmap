package sortedmap

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type NoLockSortedMapCalc[K constraints.Ordered, V any] struct {
	keys    []K
	values  []V
	calcKey func(V) K
}

func NewNoLockSortedMapCalc[K constraints.Ordered, V any](capacity int, calcKey func(V) K) *NoLockSortedMapCalc[K, V] {
	return &NoLockSortedMapCalc[K, V]{
		keys:    make([]K, 0, capacity),
		values:  make([]V, 0, capacity),
		calcKey: calcKey,
	}
}

func (s *NoLockSortedMapCalc[K, V]) Size() int {
	return len(s.values)
}

func (s *NoLockSortedMapCalc[K, V]) Capacity() int {
	return cap(s.values)
}

func (s *NoLockSortedMapCalc[K, V]) ExtendCapacityTo(newCap int) {
	if s.Capacity() < newCap {
		s.keys = append(make([]K, 0, newCap), s.keys...)
		s.values = append(make([]V, 0, newCap), s.values...)
	}
}

func (s *NoLockSortedMapCalc[K, V]) Clear() {
	s.keys = s.keys[:0]
	s.values = s.values[:0]
}

func (s *NoLockSortedMapCalc[K, V]) Insert(value V) int {
	key := s.calcKey(value)
	pos, exists := slices.BinarySearch(s.keys, key)
	if exists {
		return -1
	}

	s.keys = insertAt(s.keys, pos, key)
	s.values = insertAt(s.values, pos, value)
	return pos
}

// TODO
// func (s *NewNoLockSortedMapCalc[K, V]) InsertWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // inserted index
// }

func (s *NoLockSortedMapCalc[K, V]) InsertWithAfterHint(value V, afterIndex int) int {
	key := s.calcKey(value)
	partialKeys := s.keys[afterIndex:]
	pos, exists := slices.BinarySearch(partialKeys, key)
	if exists {
		return -1
	}

	actualPos := afterIndex + pos
	s.keys =  insertAt(s.keys, actualPos, key)
	s.values =  insertAt(s.values, actualPos, value)
	return actualPos
}

func (s *NoLockSortedMapCalc[K, V]) Delete(value V) int {
	key := s.calcKey(value)
	pos, exists := slices.BinarySearch(s.keys, key)
	if !exists {
		return -1
	}

	s.keys = deleteAt(s.keys, pos)
	s.values = deleteAt(s.values, pos)
	return pos
}

// TODO
// func (s *NewNoLockSortedMapCalc[K, V]) DeleteWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // deleted index
// }

func (s *NoLockSortedMapCalc[K, V]) DeleteWithAfterHint(value V, afterIndex int) int {
	key := s.calcKey(value)
	partialKeys := s.keys[afterIndex:]
	pos, exists := slices.BinarySearch(partialKeys, key)
	if !exists {
		return -1
	}

	actualPos := afterIndex + pos
	s.keys =  deleteAt(s.keys, actualPos)
	s.values =  deleteAt(s.values, actualPos)
	return actualPos
}

func (s *NoLockSortedMapCalc[K, V]) InsertAll(values []V) {
	s.ExtendCapacityTo(s.Size() + len(values))

	for i := range values {
		s.Insert(values[i])
	}
}

func (s *NoLockSortedMapCalc[K, V]) InsertAllOrdered(values []V) {
	s.ExtendCapacityTo(s.Size() + len(values))

	hint := 0
	for i := range values {
		hint = s.InsertWithAfterHint(values[i], hint)
	}
}


func (s *NoLockSortedMapCalc[K, V]) DeleteAll(values []V) {
	s.ExtendCapacityTo(s.Size() + len(values))

	for i := range values {
		s.Delete(values[i])
	}
}

func (s *NoLockSortedMapCalc[K, V]) DeleteAllOrdered(values []V) {
	s.ExtendCapacityTo(s.Size() + len(values))

	hint := 0
	for i := range values {
		hint = s.DeleteWithAfterHint(values[i], hint)
	}
}

func (s *NoLockSortedMapCalc[K, V]) Contains(key K) bool {
	_, exists := slices.BinarySearch(s.keys, key)
	return exists
}

func (s *NoLockSortedMapCalc[K, V]) GetIndexOfGreater(key K) int {
	pos, exists := slices.BinarySearch(s.keys, key)
	if exists {
		pos++ // does not include multiple same values
	}
	return pos
}
func (s *NoLockSortedMapCalc[K, V]) GetIndexOfGreaterOrEqual(key K) int {
	pos, _ := slices.BinarySearch(s.keys, key)
	return pos
}

func (s *NoLockSortedMapCalc[K, V]) GetGreater(key K) []V {
	pos := s.GetIndexOfGreater(key)
	return s.values[pos:]
}
func (s *NoLockSortedMapCalc[K, V]) GetGreaterOrEqual(key K) []V {
	pos := s.GetIndexOfGreaterOrEqual(key)
	return s.values[pos:]
}
func (s *NoLockSortedMapCalc[K, V]) GetLess(key K) []V {
	pos := s.GetIndexOfGreaterOrEqual(key)
	return s.values[:pos]
}
func (s *NoLockSortedMapCalc[K, V]) GetLessOrEqual(key K) []V {
	pos := s.GetIndexOfGreater(key)
	return s.values[:pos]
}

func (s *NoLockSortedMapCalc[K, V]) GetByInclusiveRange(startKey K, endKey K) []V {
	startPos := s.GetIndexOfGreaterOrEqual(startKey)
	endPos := s.GetIndexOfGreater(endKey)
	return s.values[startPos:endPos]
}
