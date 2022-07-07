package sortedmap

import (
	"testing"

	igrmkTreeMap "github.com/igrmk/treemap/v2"
	umpcSortedMap "github.com/umpc/go-sortedmap"
	umpcSortedMapAsc "github.com/umpc/go-sortedmap/asc"
	okAvlTree "gopkg.in/OlexiyKhokhlov/avltree.v2"
)

func BenchmarkNolockSet_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := NewNoLockSortedSet[int](3)
		set.Insert(300)
		set.Insert(500)
		set.Insert(700)

		set.Delete(300)
		set.Delete(500)
		set.Delete(700)
	}
}

func BenchmarkSet_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := NewSortedSet[int](3)

		set.Insert(300)
		set.Insert(500)
		set.Insert(700)

		set.Delete(300)
		set.Delete(500)
		set.Delete(700)
	}
}

func BenchmarkUmpcSortedMapSet_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := umpcSortedMap.New(3, umpcSortedMapAsc.Int)

		set.Insert(300, 300)
		set.Insert(500, 500)
		set.Insert(700, 700)

		set.Delete(300)
		set.Delete(500)
		set.Delete(700)
	}
}

func BenchmarkIgrmkTreeMapSet_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := igrmkTreeMap.New[int, struct{}]()

		set.Set(300, struct{}{})
		set.Set(500, struct{}{})
		set.Set(700, struct{}{})

		set.Del(300)
		set.Del(500)
		set.Del(700)
	}
}

func BenchmarkOkAvlTreeSet_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := okAvlTree.NewAVLTreeOrderedKey[int, struct{}]()

		set.Insert(300, struct{}{})
		set.Insert(500, struct{}{})
		set.Insert(700, struct{}{})

		set.Erase(300)
		set.Erase(500)
		set.Erase(700)
	}
}
