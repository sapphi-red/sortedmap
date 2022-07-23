package sortedmap

import (
	"strconv"
	"testing"

	igrmkTreeMap "github.com/igrmk/treemap/v2"
	okAvlTree "gopkg.in/OlexiyKhokhlov/avltree.v2"
)

const MapInsertMultipleSize = 100

func BenchmarkNoLockMap_InsertMultiple(b *testing.B) {
	var keys = make([]int, MapInsertMultipleSize)
	var values = make([]string, MapInsertMultipleSize)
	for i := range values {
		keys[i] = i * 5
		values[i] = strconv.Itoa(i * 5)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := NewNoLockSortedMap[int, string](3)

		m.InsertAll(keys, values)
	}
}

func BenchmarkNoLockMap_InsertMultipleByMap(b *testing.B) {
	var valueMap = make(map[int]string, MapInsertMultipleSize)
	for i := 0; i < MapInsertMultipleSize; i++ {
		valueMap[i*5] = strconv.Itoa(i * 5)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := NewNoLockSortedMap[int, string](3)

		m.InsertAllByMap(valueMap)
	}
}

func BenchmarkNoLockMap_InsertMultipleOrdered(b *testing.B) {
	var keys = make([]int, MapInsertMultipleSize)
	var values = make([]string, MapInsertMultipleSize)
	for i := range values {
		keys[i] = i * 5
		values[i] = strconv.Itoa(i * 5)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := NewNoLockSortedMap[int, string](3)

		m.InsertAllOrdered(keys, values)
	}
}

func BenchmarkMap_InsertMultiple(b *testing.B) {
	var keys = make([]int, MapInsertMultipleSize)
	var values = make([]string, MapInsertMultipleSize)
	for i := range values {
		keys[i] = i * 5
		values[i] = strconv.Itoa(i * 5)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := NewSortedMap[int, string](3)

		m.InsertAll(keys, values)
	}
}

func BenchmarkMap_InsertMultipleByMap(b *testing.B) {
	var valueMap = make(map[int]string, MapInsertMultipleSize)
	for i := 0; i < MapInsertMultipleSize; i++ {
		valueMap[i*5] = strconv.Itoa(i * 5)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := NewSortedMap[int, string](3)

		m.InsertAllByMap(valueMap)
	}
}

func BenchmarkNoLockMapCalc_InsertMultiple(b *testing.B) {
	var values = make([]string, MapInsertMultipleSize)
	for i := range values {
		values[i] = strconv.Itoa(i * 5)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := NewNoLockSortedMapCalc(3, safeAtoi)

		m.InsertAll(values)
	}
}

func BenchmarkMapCalc_InsertMultiple(b *testing.B) {
	var values = make([]string, MapInsertMultipleSize)
	for i := range values {
		values[i] = strconv.Itoa(i * 5)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := NewSortedMapCalc(3, safeAtoi)

		m.InsertAll(values)
	}
}

func BenchmarkIgrmkTreeMapMap_InsertMultiple(b *testing.B) {
	var keys = make([]int, MapInsertMultipleSize)
	var values = make([]string, MapInsertMultipleSize)
	for i := range values {
		keys[i] = i * 5
		values[i] = strconv.Itoa(i * 5)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := igrmkTreeMap.New[int, string]()

		for j := range keys {
			m.Set(keys[j], values[j])
		}
	}
}

func BenchmarkOkAvlTreeMap_InsertMultiple(b *testing.B) {
	var keys = make([]int, MapInsertMultipleSize)
	var values = make([]string, MapInsertMultipleSize)
	for i := range values {
		keys[i] = i * 5
		values[i] = strconv.Itoa(i * 5)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := okAvlTree.NewAVLTreeOrderedKey[int, string]()

		for j := range keys {
			m.Insert(keys[j], values[j])
		}
	}
}
