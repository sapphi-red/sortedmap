package sortedmap

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type SortedMapCalc[K constraints.Ordered, V any] struct {
	s NoLockSortedMapCalc[K, V]
	m sync.RWMutex
}

func NewSortedMapCalc[K constraints.Ordered, V any](capacity int, calcKey func(V) K) *SortedMapCalc[K, V] {
	return &SortedMapCalc[K, V]{
		s: NoLockSortedMapCalc[K, V]{
			keys:    make([]K, 0, capacity),
			values:  make([]V, 0, capacity),
			calcKey: calcKey,
		},
	}
}

func (s *SortedMapCalc[K, V]) Size() int {
	s.m.RLock()
	l := s.s.Size()
	s.m.RUnlock()
	return l
}

func (s *SortedMapCalc[K, V]) Capacity() int {
	s.m.RLock()
	c := s.s.Capacity()
	s.m.RUnlock()
	return c
}

func (s *SortedMapCalc[K, V]) ExtendCapacityTo(newCap int) {
	s.m.Lock()
	s.s.ExtendCapacityTo(newCap)
	s.m.Unlock()
}

func (s *SortedMapCalc[K, V]) Clear() {
	s.m.Lock()
	s.s.Clear()
	s.m.Unlock()
}

func (s *SortedMapCalc[K, V]) Insert(value V) int {
	s.m.Lock()
	res := s.s.Insert(value)
	s.m.Unlock()
	return res
}

// TODO
// func (s *SortedMapCalc[K, V]) InsertWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // inserted index
// }

// TODO
// func (s *SortedMapCalc[K, V]) InsertWithAfterHint(value K, afterIndex int) int {
// 	return 0 // inserted index
// }

func (s *SortedMapCalc[K, V]) Delete(key K) int {
	s.m.Lock()
	res := s.s.Delete(key)
	s.m.Unlock()
	return res
}

// TODO
// func (s *SortedMapCalc[K, V]) DeleteWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // deleted index
// }

// TODO
// func (s *SortedMapCalc[K, V]) DeleteWithAfterHint(value K, afterIndex int) int {
// 	return 0 // deleted index
// }

func (s *SortedMapCalc[K, V]) InsertAll(values []V) {
	s.m.Lock()
	s.s.InsertAll(values)
	s.m.Unlock()
}

// TODO
// func (s *SortedMapCalc[K, V]) InsertAllOrdered(values []K) {
// }

func (s *SortedMapCalc[K, V]) Contains(key K) bool {
	s.m.RLock()
	res := s.s.Contains(key)
	s.m.RUnlock()
	return res
}

func (s *SortedMapCalc[K, V]) GetIndexOfGreater(key K) int {
	s.m.RLock()
	res := s.s.GetIndexOfGreater(key)
	s.m.RUnlock()
	return res
}
func (s *SortedMapCalc[K, V]) GetIndexOfGreaterOrEqual(key K) int {
	s.m.RLock()
	res := s.s.GetIndexOfGreaterOrEqual(key)
	s.m.RUnlock()
	return res
}

func (s *SortedMapCalc[K, V]) GetGreater(key K) []V {
	s.m.RLock()
	res := s.s.GetGreater(key)
	s.m.RUnlock()
	return res
}
func (s *SortedMapCalc[K, V]) GetGreaterOrEqual(key K) []V {
	s.m.RLock()
	res := s.s.GetGreaterOrEqual(key)
	s.m.RUnlock()
	return res
}
func (s *SortedMapCalc[K, V]) GetLess(key K) []V {
	s.m.RLock()
	res := s.s.GetLess(key)
	s.m.RUnlock()
	return res
}
func (s *SortedMapCalc[K, V]) GetLessOrEqual(key K) []V {
	s.m.RLock()
	res := s.s.GetLessOrEqual(key)
	s.m.RUnlock()
	return res
}

func (s *SortedMapCalc[K, V]) GetByInclusiveRange(startKey K, endKey K) []V {
	s.m.RLock()
	res := s.s.GetByInclusiveRange(startKey, endKey)
	s.m.RUnlock()
	return res
}
