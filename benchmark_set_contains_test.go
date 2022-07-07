package sortedmap

import (
	"testing"

	igrmkTreeMap "github.com/igrmk/treemap/v2"
	umpcSortedMap "github.com/umpc/go-sortedmap"
	umpcSortedMapAsc "github.com/umpc/go-sortedmap/asc"
	okAvlTree "gopkg.in/OlexiyKhokhlov/avltree.v2"
)

const SetContainsSize = 10000

func BenchmarkNoLockSet_Contains(b *testing.B) {
	set := NewNoLockSortedSet[int](SetContainsSize)
	for i := 0; i < set.Capacity(); i++ {
		set.Insert(i * 3)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkSContains = set.Contains(300)
		sinkSContains = set.Contains(500)
		sinkSContains = set.Contains(700)
	}
}

func BenchmarkSet_Contains(b *testing.B) {
	set := NewSortedSet[int](SetContainsSize)
	for i := 0; i < set.Capacity(); i++ {
		set.Insert(i * 3)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkSContains = set.Contains(300)
		sinkSContains = set.Contains(500)
		sinkSContains = set.Contains(700)
	}
}

func BenchmarkUmpcSortedMapSet_Contains(b *testing.B) {
	set := umpcSortedMap.New(SetContainsSize, umpcSortedMapAsc.Int)
	for i := 0; i < SetContainsSize; i++ {
		set.Insert(i*3, i*3)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, sinkSContains = set.Get(300)
		_, sinkSContains = set.Get(500)
		_, sinkSContains = set.Get(700)
	}
}

func BenchmarkIgrmkTreeMapSet_Contains(b *testing.B) {
	set := igrmkTreeMap.New[int, struct{}]()
	for i := 0; i < SetContainsSize; i++ {
		set.Set(i*3, struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, sinkSContains = set.Get(300)
		_, sinkSContains = set.Get(500)
		_, sinkSContains = set.Get(700)
	}
}

func BenchmarkOkAvlTreeSet_Contains(b *testing.B) {
	set := okAvlTree.NewAVLTreeOrderedKey[int, struct{}]()
	for i := 0; i < SetContainsSize; i++ {
		set.Insert(i*3, struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkSContains = set.Find(300) != nil
		sinkSContains = set.Find(500) != nil
		sinkSContains = set.Find(700) != nil
	}
}

var sinkSContains = false
