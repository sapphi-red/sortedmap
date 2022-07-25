package sortedmap

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type SortedSet[K constraints.Ordered] struct {
	s NoLockSortedSet[K]
	m sync.RWMutex
}

func NewSortedSet[K constraints.Ordered](capacity int) *SortedSet[K] {
	return &SortedSet[K]{
		s: NoLockSortedSet[K]{
			values: make([]K, 0, capacity),
		},
	}
}

func (s *SortedSet[K]) Size() int {
	s.m.RLock()
	l := s.s.Size()
	s.m.RUnlock()
	return l
}

func (s *SortedSet[K]) Capacity() int {
	s.m.RLock()
	c := s.s.Capacity()
	s.m.RUnlock()
	return c
}

func (s *SortedSet[K]) ExtendCapacityTo(newCap int) {
	s.m.Lock()
	s.s.ExtendCapacityTo(newCap)
	s.m.Unlock()
}

func (s *SortedSet[K]) Clear() {
	s.m.Lock()
	s.s.Clear()
	s.m.Unlock()
}

func (s *SortedSet[K]) Insert(value K) int {
	s.m.Lock()
	res := s.s.Insert(value)
	s.m.Unlock()
	return res
}

// TODO
// func (s *SortedSet[K]) InsertWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // inserted index
// }

func (s *SortedSet[K]) InsertWithAfterHint(value K, afterIndex int) int {
	s.m.Lock()
	res := s.s.InsertWithAfterHint(value, afterIndex)
	s.m.Unlock()
	return res
}

func (s *SortedSet[K]) Delete(value K) int {
	s.m.Lock()
	res := s.s.Delete(value)
	s.m.Unlock()
	return res
}

// TODO
// func (s *SortedSet[K]) DeleteWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // deleted index
// }

// TODO
// func (s *SortedSet[K]) DeleteWithAfterHint(value K, afterIndex int) int {
// 	return 0 // deleted index
// }

func (s *SortedSet[K]) InsertAll(values []K) {
	s.m.Lock()
	s.s.InsertAll(values)
	s.m.Unlock()
}

func (s *SortedSet[K]) InsertAllOrdered(values []K) {
	s.m.Lock()
	s.s.InsertAllOrdered(values)
	s.m.Unlock()
}

func (s *SortedSet[K]) Contains(value K) bool {
	s.m.RLock()
	res := s.s.Contains(value)
	s.m.RUnlock()
	return res
}

func (s *SortedSet[K]) GetIndexOfGreater(value K) int {
	s.m.RLock()
	res := s.s.GetIndexOfGreater(value)
	s.m.RUnlock()
	return res
}
func (s *SortedSet[K]) GetIndexOfGreaterOrEqual(value K) int {
	s.m.RLock()
	res := s.s.GetIndexOfGreaterOrEqual(value)
	s.m.RUnlock()
	return res
}

func (s *SortedSet[K]) GetGreater(value K) []K {
	s.m.RLock()
	res := s.s.GetGreater(value)
	s.m.RUnlock()
	return res
}
func (s *SortedSet[K]) GetGreaterOrEqual(value K) []K {
	s.m.RLock()
	res := s.s.GetGreaterOrEqual(value)
	s.m.RUnlock()
	return res
}
func (s *SortedSet[K]) GetLess(value K) []K {
	s.m.RLock()
	res := s.s.GetLess(value)
	s.m.RUnlock()
	return res
}
func (s *SortedSet[K]) GetLessOrEqual(value K) []K {
	s.m.RLock()
	res := s.s.GetLessOrEqual(value)
	s.m.RUnlock()
	return res
}

func (s *SortedSet[K]) GetByInclusiveRange(startValue K, endValue K) []K {
	s.m.RLock()
	res := s.s.GetByInclusiveRange(startValue, endValue)
	s.m.RUnlock()
	return res
}
