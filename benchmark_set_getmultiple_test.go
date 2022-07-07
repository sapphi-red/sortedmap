package sortedmap

import (
	"testing"

	igrmkTreeMap "github.com/igrmk/treemap/v2"
	umpcSortedMap "github.com/umpc/go-sortedmap"
	umpcSortedMapAsc "github.com/umpc/go-sortedmap/asc"
	okAvlTree "gopkg.in/OlexiyKhokhlov/avltree.v2"
)

const SetGetMultipleSize = 100

func BenchmarkNolockSet_GetMultiple(b *testing.B) {
	set := NewNoLockSortedSet[int](SetGetMultipleSize)
	for i := 0; i < set.Capacity(); i++ {
		set.Insert(i * 3)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkGetMultiple = set.GetByInclusiveRange(24, 90)
	}
}

func BenchmarkSet_GetMultiple(b *testing.B) {
	set := NewSortedSet[int](SetGetMultipleSize)
	for i := 0; i < set.Capacity(); i++ {
		set.Insert(i * 3)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkGetMultiple = set.GetByInclusiveRange(24, 90)
	}
}

func BenchmarkUmpcSortedMapSet_GetMultiple(b *testing.B) {
	set := umpcSortedMap.New(SetGetMultipleSize, umpcSortedMapAsc.Int)
	for i := 0; i < SetGetMultipleSize; i++ {
		set.Insert(i*3, i*3)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkGetMultiple2, _ = set.BoundedKeys(24, 90)
	}
}

func BenchmarkIgrmkTreeMapSet_GetMultiple(b *testing.B) {
	set := igrmkTreeMap.New[int, struct{}]()
	for i := 0; i < SetGetMultipleSize; i++ {
		set.Set(i*3, struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := make([]int, 0)
		for iter, end := set.Range(24, 90); iter != end; iter.Next() {
			res = append(res, iter.Key())
		}
		sinkGetMultiple = res
	}
}

func BenchmarkOkAvlTreeSet_GetMultiple(b *testing.B) {
	set := okAvlTree.NewAVLTreeOrderedKey[int, struct{}]()
	for i := 0; i < SetGetMultipleSize; i++ {
		set.Insert(i*3, struct{}{})
	}
	start := 24
	end := 90

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := make([]int, 0)
		set.EnumerateDiapason(&start, &end, okAvlTree.ASCENDING, func(key int, value struct{}) bool {
			res = append(res, key)
			return true
		})
		sinkGetMultiple = res
	}
}

var sinkGetMultiple []int
var sinkGetMultiple2 []interface{}
