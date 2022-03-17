package algorithm

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

var sortedNumbers = []int{1, 2, 3, 4, 5, 6, 7, 8}

func TestLinearSearch(t *testing.T) {
	asssert := internal.NewAssert(t, "TestLinearSearch")

	comparator := &intComparator{}
	asssert.Equal(4, LinearSearch(sortedNumbers, 5, comparator))
	asssert.Equal(-1, LinearSearch(sortedNumbers, 9, comparator))
}

func TestBinarySearch(t *testing.T) {
	asssert := internal.NewAssert(t, "TestBinarySearch")

	comparator := &intComparator{}
	asssert.Equal(4, BinarySearch(sortedNumbers, 5, 0, len(sortedNumbers)-1, comparator))
	asssert.Equal(-1, BinarySearch(sortedNumbers, 9, 0, len(sortedNumbers)-1, comparator))
}

func TestBinaryIterativeSearch(t *testing.T) {
	asssert := internal.NewAssert(t, "TestBinaryIterativeSearch")

	comparator := &intComparator{}
	asssert.Equal(4, BinaryIterativeSearch(sortedNumbers, 5, 0, len(sortedNumbers)-1, comparator))
	asssert.Equal(-1, BinaryIterativeSearch(sortedNumbers, 9, 0, len(sortedNumbers)-1, comparator))
}
