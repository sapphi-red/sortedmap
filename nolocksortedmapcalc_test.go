package sortedmap_test

import (
	"testing"

	"github.com/sapphi-red/sortedmap"
	"github.com/stretchr/testify/assert"
)

func TestNoLockSortedMapCalc_Size(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	assert.Equal(t, 0, set.Size())

	set.Insert("0")
	set.Insert("3")
	assert.Equal(t, 2, set.Size())
}

func TestNoLockSortedMapCalc_Capacity(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	assert.Equal(t, 5, set.Capacity())

	set.ExtendCapacityTo(8)
	assert.Equal(t, 8, set.Capacity())
}

func TestNoLockSortedMapCalc_Clear(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	set.Insert("1")
	set.Insert("2")
	assert.Equal(t, 2, set.Size())

	set.Clear()
	assert.Equal(t, 0, set.Size())
}

func TestNoLockSortedMapCalc_Insert(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	assert.Equal(t, 0, set.Size())

	res := set.Insert("1")
	assert.Equal(t, 0, res)
	assert.Equal(t, 1, set.Size())
	assert.Equal(t, true, set.Contains(1))

	res2 := set.Insert("1")
	assert.Equal(t, -1, res2)
	assert.Equal(t, 1, set.Size())
	assert.Equal(t, true, set.Contains(1))
}

func TestNoLockSortedMapCalc_Delete(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	set.Insert("1")
	assert.Equal(t, 1, set.Size())

	res := set.Delete(1)
	assert.Equal(t, 0, res)
	assert.Equal(t, 0, set.Size())

	res2 := set.Delete(1)
	assert.Equal(t, -1, res2)
	assert.Equal(t, 0, set.Size())
	assert.Equal(t, false, set.Contains(1))
}

func TestNoLockSortedMapCalc_InsertAll(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	set.InsertAll([]string{"1", "3", "4"})
	assert.Equal(t, 3, set.Size())
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, false, set.Contains(2))
	assert.Equal(t, true, set.Contains(3))
	assert.Equal(t, true, set.Contains(4))
	assert.Equal(t, false, set.Contains(5))
	assert.Equal(t, false, set.Contains(6))

	set.InsertAll([]string{"2", "5"})
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, true, set.Contains(2))
	assert.Equal(t, true, set.Contains(3))
	assert.Equal(t, true, set.Contains(4))
	assert.Equal(t, true, set.Contains(5))
	assert.Equal(t, false, set.Contains(6))
}

func TestNoLockSortedMapCalc_Contains(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	assert.Equal(t, false, set.Contains(0))
	assert.Equal(t, false, set.Contains(1))

	set.Insert("1")
	assert.Equal(t, false, set.Contains(0))
	assert.Equal(t, true, set.Contains(1))
}

func TestNoLockSortedMapCalc_GetIndexOfGreater(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	assert.Equal(t, 0, set.GetIndexOfGreater(0))
	assert.Equal(t, 0, set.GetIndexOfGreater(3))

	set.Insert("3")
	assert.Equal(t, 0, set.GetIndexOfGreater(0))
	assert.Equal(t, 0, set.GetIndexOfGreater(2))
	assert.Equal(t, 1, set.GetIndexOfGreater(3))
	assert.Equal(t, 1, set.GetIndexOfGreater(4))
}

func TestNoLockSortedMapCalc_GetIndexOfGreaterOrEqual(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(0))
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(3))

	set.Insert("3")
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(0))
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(2))
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(3))
	assert.Equal(t, 1, set.GetIndexOfGreaterOrEqual(4))
}

func TestNoLockSortedMapCalc_GetGreater(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	assert.Equal(t, []string{}, set.GetGreater(0))

	set.Insert("3")
	assert.Equal(t, []string{"3"}, set.GetGreater(0))
	assert.Equal(t, []string{"3"}, set.GetGreater(2))
	assert.Equal(t, []string{}, set.GetGreater(3))
	assert.Equal(t, []string{}, set.GetGreater(4))
}

func TestNoLockSortedMapCalc_GetGreaterOrEqual(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	assert.Equal(t, []string{}, set.GetGreaterOrEqual(0))

	set.Insert("3")
	assert.Equal(t, []string{"3"}, set.GetGreaterOrEqual(0))
	assert.Equal(t, []string{"3"}, set.GetGreaterOrEqual(2))
	assert.Equal(t, []string{"3"}, set.GetGreaterOrEqual(3))
	assert.Equal(t, []string{}, set.GetGreaterOrEqual(4))
}

func TestNoLockSortedMapCalc_GetLess(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	assert.Equal(t, []string{}, set.GetLess(0))

	set.Insert("3")
	assert.Equal(t, []string{}, set.GetLess(0))
	assert.Equal(t, []string{}, set.GetLess(2))
	assert.Equal(t, []string{}, set.GetLess(3))
	assert.Equal(t, []string{"3"}, set.GetLess(4))
}

func TestNoLockSortedMapCalc_GetLessOrEqual(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	assert.Equal(t, []string{}, set.GetLessOrEqual(0))

	set.Insert("3")
	assert.Equal(t, []string{}, set.GetLessOrEqual(0))
	assert.Equal(t, []string{}, set.GetLessOrEqual(2))
	assert.Equal(t, []string{"3"}, set.GetLessOrEqual(3))
	assert.Equal(t, []string{"3"}, set.GetLessOrEqual(4))
}

func TestNoLockSortedMapCalc_GetByInclusiveRange(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewNoLockSortedMapCalc(5, safeAtoi)
	assert.Equal(t, []string{}, set.GetByInclusiveRange(0, 5))

	set.Insert("3")
	assert.Equal(t, []string{}, set.GetByInclusiveRange(0, 0))
	assert.Equal(t, []string{}, set.GetByInclusiveRange(0, 2))
	assert.Equal(t, []string{"3"}, set.GetByInclusiveRange(0, 3))
	assert.Equal(t, []string{"3"}, set.GetByInclusiveRange(0, 4))
	assert.Equal(t, []string{"3"}, set.GetByInclusiveRange(2, 3))
	assert.Equal(t, []string{"3"}, set.GetByInclusiveRange(2, 4))
	assert.Equal(t, []string{"3"}, set.GetByInclusiveRange(3, 3))
	assert.Equal(t, []string{"3"}, set.GetByInclusiveRange(3, 4))
	assert.Equal(t, []string{}, set.GetByInclusiveRange(4, 5))

	assert.Equal(t, []string{}, set.GetByInclusiveRange(3, 2))
	assert.Equal(t, []string{}, set.GetByInclusiveRange(5, 4))
}
