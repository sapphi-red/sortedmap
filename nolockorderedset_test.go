package orderedmap_test

import (
	"testing"

	"github.com/sapphi-red/orderedmap"
	"github.com/stretchr/testify/assert"
)

func TestNoLockOrderedSet_Size(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	assert.Equal(t, 0, set.Size())

	set.Insert(0)
	set.Insert(3)
	assert.Equal(t, 2, set.Size())
}

func TestNoLockOrderedSet_Capacity(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	assert.Equal(t, 5, set.Capacity())

	set.ExtendCapacityTo(8)
	assert.Equal(t, 8, set.Capacity())
}

func TestNoLockOrderedSet_Clear(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	set.Insert(1)
	set.Insert(2)
	assert.Equal(t, 2, set.Size())

	set.Clear()
	assert.Equal(t, 0, set.Size())
}

func TestNoLockOrderedSet_Insert(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	assert.Equal(t, 0, set.Size())

	res := set.Insert(1)
	assert.Equal(t, 0, res)
	assert.Equal(t, 1, set.Size())
	assert.Equal(t, true, set.Contains(1))

	res2 := set.Insert(1)
	assert.Equal(t, -1, res2)
	assert.Equal(t, 1, set.Size())
	assert.Equal(t, true, set.Contains(1))
}

func TestNoLockOrderedSet_Delete(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	set.Insert(1)
	assert.Equal(t, 1, set.Size())

	res := set.Delete(1)
	assert.Equal(t, 0, res)
	assert.Equal(t, 0, set.Size())

	res2 := set.Delete(1)
	assert.Equal(t, -1, res2)
	assert.Equal(t, 0, set.Size())
	assert.Equal(t, false, set.Contains(1))
}

func TestNoLockOrderedSet_InsertAll(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	set.InsertAll([]int{1, 3, 4})
	assert.Equal(t, 3, set.Size())
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, false, set.Contains(2))
	assert.Equal(t, true, set.Contains(3))
	assert.Equal(t, true, set.Contains(4))
	assert.Equal(t, false, set.Contains(5))
	assert.Equal(t, false, set.Contains(6))

	set.InsertAll([]int{2, 5})
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, true, set.Contains(2))
	assert.Equal(t, true, set.Contains(3))
	assert.Equal(t, true, set.Contains(4))
	assert.Equal(t, true, set.Contains(5))
	assert.Equal(t, false, set.Contains(6))
}

func TestNoLockOrderedSet_Contains(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	assert.Equal(t, false, set.Contains(0))
	assert.Equal(t, false, set.Contains(1))

	set.Insert(1)
	assert.Equal(t, false, set.Contains(0))
	assert.Equal(t, true, set.Contains(1))
}

func TestNoLockOrderedSet_GetIndexOfGreater(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	assert.Equal(t, 0, set.GetIndexOfGreater(0))
	assert.Equal(t, 0, set.GetIndexOfGreater(3))

	set.Insert(3)
	assert.Equal(t, 0, set.GetIndexOfGreater(0))
	assert.Equal(t, 0, set.GetIndexOfGreater(2))
	assert.Equal(t, 1, set.GetIndexOfGreater(3))
	assert.Equal(t, 1, set.GetIndexOfGreater(4))
}

func TestNoLockOrderedSet_GetIndexOfGreaterOrEqual(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(0))
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(3))

	set.Insert(3)
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(0))
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(2))
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(3))
	assert.Equal(t, 1, set.GetIndexOfGreaterOrEqual(4))
}

func TestNoLockOrderedSet_GetGreater(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	assert.Equal(t, []int{}, set.GetGreater(0))

	set.Insert(3)
	assert.Equal(t, []int{3}, set.GetGreater(0))
	assert.Equal(t, []int{3}, set.GetGreater(2))
	assert.Equal(t, []int{}, set.GetGreater(3))
	assert.Equal(t, []int{}, set.GetGreater(4))
}

func TestNoLockOrderedSet_GetGreaterOrEqual(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	assert.Equal(t, []int{}, set.GetGreaterOrEqual(0))

	set.Insert(3)
	assert.Equal(t, []int{3}, set.GetGreaterOrEqual(0))
	assert.Equal(t, []int{3}, set.GetGreaterOrEqual(2))
	assert.Equal(t, []int{3}, set.GetGreaterOrEqual(3))
	assert.Equal(t, []int{}, set.GetGreaterOrEqual(4))
}

func TestNoLockOrderedSet_GetLess(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	assert.Equal(t, []int{}, set.GetLess(0))

	set.Insert(3)
	assert.Equal(t, []int{}, set.GetLess(0))
	assert.Equal(t, []int{}, set.GetLess(2))
	assert.Equal(t, []int{}, set.GetLess(3))
	assert.Equal(t, []int{3}, set.GetLess(4))
}

func TestNoLockOrderedSet_GetLessOrEqual(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	assert.Equal(t, []int{}, set.GetLessOrEqual(0))

	set.Insert(3)
	assert.Equal(t, []int{}, set.GetLessOrEqual(0))
	assert.Equal(t, []int{}, set.GetLessOrEqual(2))
	assert.Equal(t, []int{3}, set.GetLessOrEqual(3))
	assert.Equal(t, []int{3}, set.GetLessOrEqual(4))
}

func TestNoLockOrderedSet_GetByInclusiveRange(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewNoLockOrderedSet[int](5)
	assert.Equal(t, []int{}, set.GetByInclusiveRange(0, 5))

	set.Insert(3)
	assert.Equal(t, []int{}, set.GetByInclusiveRange(0, 0))
	assert.Equal(t, []int{}, set.GetByInclusiveRange(0, 2))
	assert.Equal(t, []int{3}, set.GetByInclusiveRange(0, 3))
	assert.Equal(t, []int{3}, set.GetByInclusiveRange(0, 4))
	assert.Equal(t, []int{3}, set.GetByInclusiveRange(2, 3))
	assert.Equal(t, []int{3}, set.GetByInclusiveRange(2, 4))
	assert.Equal(t, []int{3}, set.GetByInclusiveRange(3, 3))
	assert.Equal(t, []int{3}, set.GetByInclusiveRange(3, 4))
	assert.Equal(t, []int{}, set.GetByInclusiveRange(4, 5))

	assert.Equal(t, []int{}, set.GetByInclusiveRange(3, 2))
	assert.Equal(t, []int{}, set.GetByInclusiveRange(5, 4))
}
