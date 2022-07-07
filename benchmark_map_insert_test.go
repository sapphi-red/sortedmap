package sortedmap

import (
	"testing"

	igrmkTreeMap "github.com/igrmk/treemap/v2"
	okAvlTree "gopkg.in/OlexiyKhokhlov/avltree.v2"
)

func BenchmarkNoLockMap_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := NewNoLockSortedMap[int, string](3)

		m.Insert(300, "300")
		m.Insert(500, "500")
		m.Insert(700, "700")
	}
}

func BenchmarkMap_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := NewSortedMap[int, string](3)

		m.Insert(300, "300")
		m.Insert(500, "500")
		m.Insert(700, "700")
	}
}

func BenchmarkNoLockMapCalc_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := NewNoLockSortedMapCalc(3, safeAtoi)

		m.Insert("300")
		m.Insert("500")
		m.Insert("700")
	}
}

func BenchmarkMapCalc_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := NewSortedMapCalc(3, safeAtoi)

		m.Insert("300")
		m.Insert("500")
		m.Insert("700")
	}
}

func BenchmarkIgrmkTreeMapMap_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := igrmkTreeMap.New[int, string]()

		m.Set(300, "300")
		m.Set(500, "500")
		m.Set(700, "700")
	}
}

func BenchmarkOkAvlTreeMap_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := okAvlTree.NewAVLTreeOrderedKey[int, string]()

		m.Insert(300, "300")
		m.Insert(500, "500")
		m.Insert(700, "700")
	}
}

func BenchmarkNolockMap_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewNoLockSortedMap[int, string](3)
	}
}

func BenchmarkMap_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSortedMap[int, string](3)
	}
}

func BenchmarkNolockMapCalc_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewNoLockSortedMapCalc(3, safeAtoi)
	}
}

func BenchmarkMapCalc_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSortedMapCalc(3, safeAtoi)
	}
}

func BenchmarkIgrmkTreeMapMap_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		igrmkTreeMap.New[int, string]()
	}
}

func BenchmarkOkAvlTreeMap_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		okAvlTree.NewAVLTreeOrderedKey[int, string]()
	}
}
