package sortedmap

import (
	"testing"

	igrmkTreeMap "github.com/igrmk/treemap/v2"
	umpcSortedMap "github.com/umpc/go-sortedmap"
	umpcSortedMapAsc "github.com/umpc/go-sortedmap/asc"
	okAvlTree "gopkg.in/OlexiyKhokhlov/avltree.v2"
)

const SetInsertMultipleSize = 100

func BenchmarkNolockSet_InsertMultiple(b *testing.B) {
	var values = make([]int, SetInsertMultipleSize)
	for i := range values {
		values[i] = i * 5
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set := NewNoLockSortedSet[int](3)

		set.InsertAll(values)
	}
}

func BenchmarkSet_InsertMultiple(b *testing.B) {
	var values = make([]int, SetInsertMultipleSize)
	for i := range values {
		values[i] = i * 5
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set := NewSortedSet[int](3)

		set.InsertAll(values)
	}
}

func BenchmarkUmpcSortedMapSet_InsertMultiple(b *testing.B) {
	var records = make([]umpcSortedMap.Record, SetInsertMultipleSize)
	for i := range records {
		records[i] = umpcSortedMap.Record{
			Key: i * 5,
			Val: i * 5,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set := umpcSortedMap.New(3, umpcSortedMapAsc.Int)

		set.BatchInsert(records)
	}
}

func BenchmarkIgrmkTreeMapSet_InsertMultiple(b *testing.B) {
	var values = make([]int, SetInsertMultipleSize)
	for i := range values {
		values[i] = i * 5
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set := igrmkTreeMap.New[int, struct{}]()

		for j := range values {
			set.Set(values[j], struct{}{})
		}
	}
}

func BenchmarkOkAvlTreeSet_InsertMultiple(b *testing.B) {
	var values = make([]int, SetInsertMultipleSize)
	for i := range values {
		values[i] = i * 5
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set := okAvlTree.NewAVLTreeOrderedKey[int, struct{}]()

		for j := range values {
			set.Insert(values[j], struct{}{})
		}
	}
}
