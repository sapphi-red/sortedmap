package sortedmap

import (
	"testing"

	igrmkTreeMap "github.com/igrmk/treemap/v2"
	okAvlTree "gopkg.in/OlexiyKhokhlov/avltree.v2"
)

func BenchmarkNoLockMap_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := NewNoLockSortedMap[int, string](3)
		m.Insert(300, "300")
		m.Insert(500, "500")
		m.Insert(700, "700")

		m.Delete(300)
		m.Delete(500)
		m.Delete(700)
	}
}

func BenchmarkMap_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := NewSortedMap[int, string](3)

		m.Insert(300, "300")
		m.Insert(500, "500")
		m.Insert(700, "700")

		m.Delete(300)
		m.Delete(500)
		m.Delete(700)
	}
}

func BenchmarkNoLockMapCalc_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := NewNoLockSortedMapCalc(3, safeAtoi)
		m.Insert("300")
		m.Insert("500")
		m.Insert("700")

		m.Delete(300)
		m.Delete(500)
		m.Delete(700)
	}
}

func BenchmarkMapCalc_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := NewSortedMapCalc(3, safeAtoi)
		m.Insert("300")
		m.Insert("500")
		m.Insert("700")

		m.Delete(300)
		m.Delete(500)
		m.Delete(700)
	}
}

func BenchmarkIgrmkTreeMapMap_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := igrmkTreeMap.New[int, string]()

		m.Set(300, "300")
		m.Set(500, "500")
		m.Set(700, "700")

		m.Del(300)
		m.Del(500)
		m.Del(700)
	}
}

func BenchmarkOkAvlTreeMap_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := okAvlTree.NewAVLTreeOrderedKey[int, string]()

		m.Insert(300, "300")
		m.Insert(500, "500")
		m.Insert(700, "700")

		m.Erase(300)
		m.Erase(500)
		m.Erase(700)
	}
}
