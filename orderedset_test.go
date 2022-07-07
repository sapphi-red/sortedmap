package orderedmap_test

import (
	"testing"

	"github.com/sapphi-red/orderedmap"
	"github.com/stretchr/testify/assert"
)

func TestOrderedSet_Size(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.Insert(0)
	set.Insert(3)
	assert.Equal(t, 2, set.Size())
}

func TestOrderedSet_Capacity(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.ExtendCapacityTo(8)
	assert.Equal(t, 8, set.Capacity())
}

func TestOrderedSet_Clear(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.Insert(1)
	set.Clear()
	assert.Equal(t, 0, set.Size())
}

func TestOrderedSet_Insert(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	res := set.Insert(1)
	assert.Equal(t, 0, res)
	assert.Equal(t, 1, set.Size())
	assert.Equal(t, true, set.Contains(1))
}

func TestOrderedSet_Delete(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.Insert(1)

	res := set.Delete(1)
	assert.Equal(t, 0, res)
	assert.Equal(t, 0, set.Size())
}

func TestOrderedSet_InsertAll(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.InsertAll([]int{1, 3, 4})
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, true, set.Contains(3))
	assert.Equal(t, true, set.Contains(4))
}

func TestOrderedSet_Contains(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.Insert(1)
	assert.Equal(t, false, set.Contains(0))
	assert.Equal(t, true, set.Contains(1))
}

func TestOrderedSet_GetIndexOfGreater(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.Insert(3)
	assert.Equal(t, 0, set.GetIndexOfGreater(2))
	assert.Equal(t, 1, set.GetIndexOfGreater(3))
	assert.Equal(t, 1, set.GetIndexOfGreater(4))
}

func TestOrderedSet_GetIndexOfGreaterOrEqual(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.Insert(3)
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(2))
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(3))
	assert.Equal(t, 1, set.GetIndexOfGreaterOrEqual(4))
}

func TestOrderedSet_GetGreater(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.Insert(3)
	assert.Equal(t, []int{3}, set.GetGreater(2))
	assert.Equal(t, []int{}, set.GetGreater(3))
	assert.Equal(t, []int{}, set.GetGreater(4))
}

func TestOrderedSet_GetGreaterOrEqual(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.Insert(3)
	assert.Equal(t, []int{3}, set.GetGreaterOrEqual(2))
	assert.Equal(t, []int{3}, set.GetGreaterOrEqual(3))
	assert.Equal(t, []int{}, set.GetGreaterOrEqual(4))
}

func TestOrderedSet_GetLess(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.Insert(3)
	assert.Equal(t, []int{}, set.GetLess(2))
	assert.Equal(t, []int{}, set.GetLess(3))
	assert.Equal(t, []int{3}, set.GetLess(4))
}

func TestOrderedSet_GetLessOrEqual(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.Insert(3)
	assert.Equal(t, []int{}, set.GetLessOrEqual(2))
	assert.Equal(t, []int{3}, set.GetLessOrEqual(3))
	assert.Equal(t, []int{3}, set.GetLessOrEqual(4))
}

func TestOrderedSet_GetByInclusiveRange(t *testing.T) {
	t.Parallel()

	set := orderedmap.NewOrderedSet[int](5)
	set.Insert(3)
	assert.Equal(t, []int{}, set.GetByInclusiveRange(0, 2))
	assert.Equal(t, []int{3}, set.GetByInclusiveRange(3, 3))
	assert.Equal(t, []int{}, set.GetByInclusiveRange(4, 5))
}
