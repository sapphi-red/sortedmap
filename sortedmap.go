package sortedmap

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type SortedMap[K constraints.Ordered, V any] struct {
	s NoLockSortedMap[K, V]
	m sync.RWMutex
}

func NewSortedMap[K constraints.Ordered, V any](capacity int) *SortedMap[K, V] {
	return &SortedMap[K, V]{
		s: NoLockSortedMap[K, V]{
			keys: make([]K, 0, capacity),
			values: make([]V, 0, capacity),
		},
	}
}

func (s *SortedMap[K, V]) Size() int {
	s.m.RLock()
	l := s.s.Size()
	s.m.RUnlock()
	return l
}

func (s *SortedMap[K, V]) Capacity() int {
	s.m.RLock()
	c := s.s.Capacity()
	s.m.RUnlock()
	return c
}

func (s *SortedMap[K, V]) ExtendCapacityTo(newCap int) {
	s.m.Lock()
	s.s.ExtendCapacityTo(newCap)
	s.m.Unlock()
}

func (s *SortedMap[K, V]) Clear() {
	s.m.Lock()
	s.s.Clear()
	s.m.Unlock()
}

func (s *SortedMap[K, V]) Insert(key K, value V) int {
	s.m.Lock()
	res := s.s.Insert(key, value)
	s.m.Unlock()
	return res
}

// TODO
// func (s *SortedMap[K, V]) InsertWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // inserted index
// }

func (s *SortedMap[K, V]) InsertWithAfterHint(key K, value V, afterIndex int) int {
	s.m.Lock()
	res := s.s.InsertWithAfterHint(key, value, afterIndex)
	s.m.Unlock()
	return res
}

func (s *SortedMap[K, V]) Delete(key K) int {
	s.m.Lock()
	res := s.s.Delete(key)
	s.m.Unlock()
	return res
}

// TODO
// func (s *SortedMap[K, V]) DeleteWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // deleted index
// }

// TODO
// func (s *SortedMap[K, V]) DeleteWithAfterHint(value K, afterIndex int) int {
// 	return 0 // deleted index
// }

func (s *SortedMap[K, V]) InsertAll(keys []K, values []V) {
	s.m.Lock()
	s.s.InsertAll(keys, values)
	s.m.Unlock()
}

func (s *SortedMap[K, V]) InsertAllByMap(m map[K]V) {
	s.m.Lock()
	s.s.InsertAllByMap(m)
	s.m.Unlock()
}

func (s *SortedMap[K, V]) InsertAllOrdered(keys []K, values []V) {
	s.m.Lock()
	s.s.InsertAllOrdered(keys, values)
	s.m.Unlock()
}

func (s *SortedMap[K, V]) Contains(key K) bool {
	s.m.RLock()
	res := s.s.Contains(key)
	s.m.RUnlock()
	return res
}

func (s *SortedMap[K, V]) GetIndexOfGreater(key K) int {
	s.m.RLock()
	res := s.s.GetIndexOfGreater(key)
	s.m.RUnlock()
	return res
}
func (s *SortedMap[K, V]) GetIndexOfGreaterOrEqual(key K) int {
	s.m.RLock()
	res := s.s.GetIndexOfGreaterOrEqual(key)
	s.m.RUnlock()
	return res
}

func (s *SortedMap[K, V]) GetGreater(key K) []V {
	s.m.RLock()
	res := s.s.GetGreater(key)
	s.m.RUnlock()
	return res
}
func (s *SortedMap[K, V]) GetGreaterOrEqual(key K) []V {
	s.m.RLock()
	res := s.s.GetGreaterOrEqual(key)
	s.m.RUnlock()
	return res
}
func (s *SortedMap[K, V]) GetLess(key K) []V {
	s.m.RLock()
	res := s.s.GetLess(key)
	s.m.RUnlock()
	return res
}
func (s *SortedMap[K, V]) GetLessOrEqual(key K) []V {
	s.m.RLock()
	res := s.s.GetLessOrEqual(key)
	s.m.RUnlock()
	return res
}

func (s *SortedMap[K, V]) GetByInclusiveRange(startKey K, endKey K) []V {
	s.m.RLock()
	res := s.s.GetByInclusiveRange(startKey, endKey)
	s.m.RUnlock()
	return res
}
