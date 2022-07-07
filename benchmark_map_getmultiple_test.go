package sortedmap

import (
	"strconv"
	"testing"

	igrmkTreeMap "github.com/igrmk/treemap/v2"
	okAvlTree "gopkg.in/OlexiyKhokhlov/avltree.v2"
)

const MapGetMultipleSize = 100

func BenchmarkNoLockMap_GetMultiple(b *testing.B) {
	m := NewNoLockSortedMap[int, string](MapGetMultipleSize)
	for i := 0; i < m.Capacity(); i++ {
		m.Insert(i*3, strconv.Itoa(i*3))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkMGetMultiple = m.GetByInclusiveRange(24, 90)
	}
}

func BenchmarkMap_GetMultiple(b *testing.B) {
	m := NewSortedMap[int, string](MapGetMultipleSize)
	for i := 0; i < m.Capacity(); i++ {
		m.Insert(i*3, strconv.Itoa(i*3))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkMGetMultiple = m.GetByInclusiveRange(24, 90)
	}
}

func BenchmarkIgrmkTreeMapMap_GetMultiple(b *testing.B) {
	m := igrmkTreeMap.New[int, string]()
	for i := 0; i < MapGetMultipleSize; i++ {
		m.Set(i*3, strconv.Itoa(i*3))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := make([]string, 0)
		for iter, end := m.Range(24, 90); iter != end; iter.Next() {
			res = append(res, iter.Value())
		}
		sinkMGetMultiple = res
	}
}

func BenchmarkOkAvlTreeMap_GetMultiple(b *testing.B) {
	m := okAvlTree.NewAVLTreeOrderedKey[int, string]()
	for i := 0; i < MapGetMultipleSize; i++ {
		m.Insert(i*3, strconv.Itoa(i*3))
	}
	start := 24
	end := 90

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := make([]string, 0)
		m.EnumerateDiapason(&start, &end, okAvlTree.ASCENDING, func(key int, value string) bool {
			res = append(res, value)
			return true
		})
		sinkMGetMultiple = res
	}
}

var sinkMGetMultiple []string
