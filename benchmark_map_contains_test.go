package sortedmap

import (
	"strconv"
	"testing"

	igrmkTreeMap "github.com/igrmk/treemap/v2"
	okAvlTree "gopkg.in/OlexiyKhokhlov/avltree.v2"
)

const MapContainsSize = 10000

func BenchmarkNoLockMap_Contains(b *testing.B) {
	m := NewNoLockSortedMap[int, string](MapContainsSize)
	for i := 0; i < m.Capacity(); i++ {
		m.Insert(i*3, strconv.Itoa(i*3))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkMContains = m.Contains(300)
		sinkMContains = m.Contains(500)
		sinkMContains = m.Contains(700)
	}
}

func BenchmarkMap_Contains(b *testing.B) {
	m := NewSortedMap[int, string](MapContainsSize)
	for i := 0; i < m.Capacity(); i++ {
		m.Insert(i*3, strconv.Itoa(i*3))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkMContains = m.Contains(300)
		sinkMContains = m.Contains(500)
		sinkMContains = m.Contains(700)
	}
}

func BenchmarkIgrmkTreeMapMap_Contains(b *testing.B) {
	m := igrmkTreeMap.New[int, string]()
	for i := 0; i < MapContainsSize; i++ {
		m.Set(i*3, strconv.Itoa(i*3))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, sinkMContains = m.Get(300)
		_, sinkMContains = m.Get(500)
		_, sinkMContains = m.Get(700)
	}
}

func BenchmarkOkAvlTreeMap_Contains(b *testing.B) {
	m := okAvlTree.NewAVLTreeOrderedKey[int, string]()
	for i := 0; i < MapContainsSize; i++ {
		m.Insert(i*3, strconv.Itoa(i*3))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkMContains = m.Find(300) != nil
		sinkMContains = m.Find(500) != nil
		sinkMContains = m.Find(700) != nil
	}
}

var sinkMContains = false
