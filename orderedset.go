package orderedmap

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type OrderedSet[K constraints.Ordered] struct {
	s NoLockOrderedSet[K]
	m sync.RWMutex
}

func NewOrderedSet[K constraints.Ordered](capacity int) *OrderedSet[K] {
	return &OrderedSet[K]{
		s: NoLockOrderedSet[K]{
			values: make([]K, 0, capacity),
		},
	}
}

func (s *OrderedSet[K]) Size() int {
	s.m.RLock()
	l := s.s.Size()
	s.m.RUnlock()
	return l
}

func (s *OrderedSet[K]) Capacity() int {
	s.m.RLock()
	c := s.s.Capacity()
	s.m.RUnlock()
	return c
}

func (s *OrderedSet[K]) ExtendCapacityTo(newCap int) {
	s.m.Lock()
	s.s.ExtendCapacityTo(newCap)
	s.m.Unlock()
}

func (s *OrderedSet[K]) Clear() {
	s.m.Lock()
	s.s.Clear()
	s.m.Unlock()
}

func (s *OrderedSet[K]) Insert(value K) int {
	s.m.Lock()
	res := s.s.Insert(value)
	s.m.Unlock()
	return res
}

// TODO
// func (s *OrderedSet[K]) InsertWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // inserted index
// }

// TODO
// func (s *OrderedSet[K]) InsertWithAfterHint(value K, afterIndex int) int {
// 	return 0 // inserted index
// }

func (s *OrderedSet[K]) Delete(value K) int {
	s.m.Lock()
	res := s.s.Delete(value)
	s.m.Unlock()
	return res
}

// TODO
// func (s *OrderedSet[K]) DeleteWithBeforeHint(value K, beforeIndex int) int {
// 	return 0 // deleted index
// }

// TODO
// func (s *OrderedSet[K]) DeleteWithAfterHint(value K, afterIndex int) int {
// 	return 0 // deleted index
// }

func (s *OrderedSet[K]) InsertAll(values []K) {
	s.m.Lock()
	s.s.InsertAll(values)
	s.m.Unlock()
}

// TODO
// func (s *OrderedSet[K]) InsertAllOrdered(values []K) {
// }

func (s *OrderedSet[K]) Contains(value K) bool {
	s.m.RLock()
	res := s.s.Contains(value)
	s.m.RUnlock()
	return res
}

func (s *OrderedSet[K]) GetIndexOfGreater(value K) int {
	s.m.RLock()
	res := s.s.GetIndexOfGreater(value)
	s.m.RUnlock()
	return res
}
func (s *OrderedSet[K]) GetIndexOfGreaterOrEqual(value K) int {
	s.m.RLock()
	res := s.s.GetIndexOfGreaterOrEqual(value)
	s.m.RUnlock()
	return res
}

func (s *OrderedSet[K]) GetGreater(value K) []K {
	s.m.RLock()
	res := s.s.GetGreater(value)
	s.m.RUnlock()
	return res
}
func (s *OrderedSet[K]) GetGreaterOrEqual(value K) []K {
	s.m.RLock()
	res := s.s.GetGreaterOrEqual(value)
	s.m.RUnlock()
	return res
}
func (s *OrderedSet[K]) GetLess(value K) []K {
	s.m.RLock()
	res := s.s.GetLess(value)
	s.m.RUnlock()
	return res
}
func (s *OrderedSet[K]) GetLessOrEqual(value K) []K {
	s.m.RLock()
	res := s.s.GetLessOrEqual(value)
	s.m.RUnlock()
	return res
}

func (s *OrderedSet[K]) GetByInclusiveRange(startValue K, endValue K) []K {
	s.m.RLock()
	res := s.s.GetByInclusiveRange(startValue, endValue)
	s.m.RUnlock()
	return res
}
