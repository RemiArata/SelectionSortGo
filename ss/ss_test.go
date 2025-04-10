package ss

import (
	"testing"
)

func TestSortShort(t *testing.T) {
	unsorted := []int{3, 1, 2}
	correct := []int{1, 2, 3}
	sorted := SelectionSort(unsorted)

	for i := 0; i < len(sorted); i++ {
		if correct[i] != sorted[i] {
			t.Error("Unmatched data at index", i)
		}
	}
}

func TestSortLong(t *testing.T) {
	unsorted := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	correct := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sorted := SelectionSort(unsorted)

	for i := 0; i < len(sorted); i++ {
		if correct[i] != sorted[i] {
			t.Error("Unmatched data at index", i)
		}
	}
}
