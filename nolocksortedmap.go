package sortedmap

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type NoLockSortedMap[K constraints.Ordered, V any] struct {
	keys   []K
	values []V
}

func NewNoLockSortedMap[K constraints.Ordered, V any](capacity int) *NoLockSortedMap[K, V] {
	return &NoLockSortedMap[K, V]{
		keys: make([]K, 0, capacity),
		values: make([]V, 0, capacity),
	}
}

func (s *NoLockSortedMap[K, V]) Size() int {
	return len(s.values)
}

func (s *NoLockSortedMap[K, V]) Capacity() int {
	return cap(s.values)
}

func (s *NoLockSortedMap[K, V]) ExtendCapacityTo(newCap int) {
	if s.Capacity() < newCap {
		s.keys = append(make([]K, 0, newCap), s.keys...)
		s.values = append(make([]V, 0, newCap), s.values...)
	}
}

func (s *NoLockSortedMap[K, V]) Clear() {
	s.keys = s.keys[:0]
	s.values = s.values[:0]
}

func (s *NoLockSortedMap[K, V]) Insert(key K, value V) int {
	pos, exists := slices.BinarySearch(s.keys, key)
	if exists {
		return -1
	}

	s.keys =  insertAt(s.keys, pos, key)
	s.values =  insertAt(s.values, pos, value)
	return pos
}

// TODO
// func (s *NewNoLockSortedMap[K, V]) InsertWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // inserted index
// }

// TODO
// func (s *NewNoLockSortedMap[K, V]) InsertWithAfterHint(value K, afterIndex int) int {
// 	return 0 // inserted index
// }

func (s *NoLockSortedMap[K, V]) Delete(key K) int {
	pos, exists := slices.BinarySearch(s.keys, key)
	if !exists {
		return -1
	}

	s.keys = deleteAt(s.keys, pos)
	s.values = deleteAt(s.values, pos)
	return pos
}

// TODO
// func (s *NewNoLockSortedMap[K, V]) DeleteWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // deleted index
// }

// TODO
// func (s *NewNoLockSortedMap[K, V]) DeleteWithAfterHint(value K, afterIndex int) int {
// 	return 0 // deleted index
// }

func (s *NoLockSortedMap[K, V]) InsertAll(keys []K, values []V) {
	s.ExtendCapacityTo(s.Size() + len(values))

	for i := range keys {
		s.Insert(keys[i], values[i])
	}
}

func (s *NoLockSortedMap[K, V]) InsertAllByMap(m map[K]V) {
	s.ExtendCapacityTo(s.Size() + len(m))

	for k, v := range m {
		s.Insert(k, v)
	}
}

// TODO
// func (s *NewNoLockSortedMap[K, V]) InsertAllOrdered(values []K) {
// }

func (s *NoLockSortedMap[K, V]) Contains(key K) bool {
	_, exists := slices.BinarySearch(s.keys, key)
	return exists
}

func (s *NoLockSortedMap[K, V]) GetIndexOfGreater(key K) int {
	pos, exists := slices.BinarySearch(s.keys, key)
	if exists {
		pos++ // does not include multiple same values
	}
	return pos
}
func (s *NoLockSortedMap[K, V]) GetIndexOfGreaterOrEqual(key K) int {
	pos, _ := slices.BinarySearch(s.keys, key)
	return pos
}

func (s *NoLockSortedMap[K, V]) GetGreater(key K) []V {
	pos := s.GetIndexOfGreater(key)
	return s.values[pos:]
}
func (s *NoLockSortedMap[K, V]) GetGreaterOrEqual(key K) []V {
	pos := s.GetIndexOfGreaterOrEqual(key)
	return s.values[pos:]
}
func (s *NoLockSortedMap[K, V]) GetLess(key K) []V {
	pos := s.GetIndexOfGreaterOrEqual(key)
	return s.values[:pos]
}
func (s *NoLockSortedMap[K, V]) GetLessOrEqual(key K) []V {
	pos := s.GetIndexOfGreater(key)
	return s.values[:pos]
}

func (s *NoLockSortedMap[K, V]) GetByInclusiveRange(startKey K, endKey K) []V {
	startPos := s.GetIndexOfGreaterOrEqual(startKey)
	endPos := s.GetIndexOfGreater(endKey)
	return s.values[startPos:endPos]
}
