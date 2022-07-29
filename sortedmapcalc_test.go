package sortedmap_test

import (
	"testing"

	"github.com/sapphi-red/sortedmap"
	"github.com/stretchr/testify/assert"
)

func TestSortedMapCalc_Size(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("0")
	set.Insert("3")
	assert.Equal(t, 2, set.Size())
}

func TestSortedMapCalc_Capacity(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.ExtendCapacityTo(8)
	assert.Equal(t, 8, set.Capacity())
}

func TestSortedMapCalc_Clear(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("1")
	set.Clear()
	assert.Equal(t, 0, set.Size())
}

func TestSortedMapCalc_Insert(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	res := set.Insert("1")
	assert.Equal(t, 0, res)
	assert.Equal(t, 1, set.Size())
	assert.Equal(t, true, set.Contains(1))
}

func TestSortedMapCalc_InsertWithAfterHint(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	res := set.InsertWithAfterHint("1", 0)
	assert.Equal(t, 0, res)
	assert.Equal(t, 1, set.Size())
	assert.Equal(t, true, set.Contains(1))
}

func TestSortedMapCalc_Delete(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("1")

	res := set.Delete("1")
	assert.Equal(t, 0, res)
	assert.Equal(t, 0, set.Size())
}

func TestSortedMapCalc_DeleteWithAfterHint(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("1")

	res := set.DeleteWithAfterHint("1", 0)
	assert.Equal(t, 0, res)
	assert.Equal(t, 0, set.Size())
}

func TestSortedMapCalc_InsertAll(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.InsertAll([]string{"1", "3", "4"})
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, true, set.Contains(3))
	assert.Equal(t, true, set.Contains(4))
}

func TestSortedMapCalc_InsertAllOrdered(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.InsertAllOrdered([]string{"1", "3", "4"})
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, true, set.Contains(3))
	assert.Equal(t, true, set.Contains(4))
}

func TestSortedMapCalc_DeleteAll(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.InsertAll([]string{"1", "3", "4"})
	set.DeleteAll([]string{"1", "3", "4"})
	assert.Equal(t, 0, set.Size())
}

func TestSortedMapCalc_DeleteAllOrdered(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.InsertAll([]string{"1", "3", "4"})
	set.DeleteAllOrdered([]string{"1", "3", "4"})
	assert.Equal(t, 0, set.Size())
}

func TestSortedMapCalc_Contains(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("1")
	assert.Equal(t, false, set.Contains(0))
	assert.Equal(t, true, set.Contains(1))
}

func TestSortedMapCalc_GetIndexOfGreater(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("3")
	assert.Equal(t, 0, set.GetIndexOfGreater(2))
	assert.Equal(t, 1, set.GetIndexOfGreater(3))
	assert.Equal(t, 1, set.GetIndexOfGreater(4))
}

func TestSortedMapCalc_GetIndexOfGreaterOrEqual(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("3")
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(2))
	assert.Equal(t, 0, set.GetIndexOfGreaterOrEqual(3))
	assert.Equal(t, 1, set.GetIndexOfGreaterOrEqual(4))
}

func TestSortedMapCalc_GetGreater(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("3")
	assert.Equal(t, []string{"3"}, set.GetGreater(2))
	assert.Equal(t, []string{}, set.GetGreater(3))
	assert.Equal(t, []string{}, set.GetGreater(4))
}

func TestSortedMapCalc_GetGreaterOrEqual(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("3")
	assert.Equal(t, []string{"3"}, set.GetGreaterOrEqual(2))
	assert.Equal(t, []string{"3"}, set.GetGreaterOrEqual(3))
	assert.Equal(t, []string{}, set.GetGreaterOrEqual(4))
}

func TestSortedMapCalc_GetLess(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("3")
	assert.Equal(t, []string{}, set.GetLess(2))
	assert.Equal(t, []string{}, set.GetLess(3))
	assert.Equal(t, []string{"3"}, set.GetLess(4))
}

func TestSortedMapCalc_GetLessOrEqual(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("3")
	assert.Equal(t, []string{}, set.GetLessOrEqual(2))
	assert.Equal(t, []string{"3"}, set.GetLessOrEqual(3))
	assert.Equal(t, []string{"3"}, set.GetLessOrEqual(4))
}

func TestSortedMapCalc_GetByInclusiveRange(t *testing.T) {
	t.Parallel()

	set := sortedmap.NewSortedMapCalc(5, safeAtoi)
	set.Insert("3")
	assert.Equal(t, []string{}, set.GetByInclusiveRange(0, 2))
	assert.Equal(t, []string{"3"}, set.GetByInclusiveRange(3, 3))
	assert.Equal(t, []string{}, set.GetByInclusiveRange(4, 5))
}
