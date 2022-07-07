package sortedmap

import (
	"testing"

	igrmkTreeMap "github.com/igrmk/treemap/v2"
	umpcSortedMap "github.com/umpc/go-sortedmap"
	umpcSortedMapAsc "github.com/umpc/go-sortedmap/asc"
	okAvlTree "gopkg.in/OlexiyKhokhlov/avltree.v2"
)

func BenchmarkNolockSet_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := NewNoLockSortedSet[int](3)

		set.Insert(300)
		set.Insert(500)
		set.Insert(700)
	}
}

func BenchmarkSet_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := NewSortedSet[int](3)

		set.Insert(300)
		set.Insert(500)
		set.Insert(700)
	}
}

func BenchmarkUmpcSortedMapSet_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := umpcSortedMap.New(3, umpcSortedMapAsc.Int)

		set.Insert(300, 300)
		set.Insert(500, 500)
		set.Insert(700, 700)
	}
}

func BenchmarkIgrmkTreeMapSet_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := igrmkTreeMap.New[int, struct{}]()

		set.Set(300, struct{}{})
		set.Set(500, struct{}{})
		set.Set(700, struct{}{})
	}
}

func BenchmarkOkAvlTreeSet_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := okAvlTree.NewAVLTreeOrderedKey[int, struct{}]()

		set.Insert(300, struct{}{})
		set.Insert(500, struct{}{})
		set.Insert(700, struct{}{})
	}
}

func BenchmarkNolockSet_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewNoLockSortedSet[int](3)
	}
}

func BenchmarkSet_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSortedSet[int](3)
	}
}

func BenchmarkUmpcSortedMapSet_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		umpcSortedMap.New(3, umpcSortedMapAsc.Int)
	}
}

func BenchmarkIgrmkTreeMapSet_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		igrmkTreeMap.New[int, struct{}]()
	}
}

func BenchmarkOkAvlTreeSet_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		okAvlTree.NewAVLTreeOrderedKey[int, struct{}]()
	}
}
