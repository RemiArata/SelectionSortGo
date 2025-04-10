package ss

import (
	"testing"
)

// [26 36 89 80 87 66 86 49 56 14]
func TestWholeListLast(t *testing.T) {
	data := []int{26, 36, 89, 80, 87, 66, 86, 49, 56, 14}
	fIdx, fVal := findMin(data, 0)

	if fIdx != 9 {
		t.Error("The found index: ", fIdx, "should be 9")
	}
	if fVal != 14 {
		t.Error("The found val: ", fVal, "should be 14")
	}
}

func TestWholeListFirst(t *testing.T) {
	data := []int{14, 36, 89, 80, 87, 66, 86, 49, 56, 26}
	fIdx, fVal := findMin(data, 0)

	if fIdx != 0 {
		t.Error("The found index: ", fIdx, "should be 9")
	}
	if fVal != 14 {
		t.Error("The found val: ", fVal, "should be 14")
	}
}

func TestWholeListMiddle(t *testing.T) {
	data := []int{87, 36, 89, 80, 14, 66, 86, 49, 56, 26}
	fIdx, fVal := findMin(data, 0)

	if fIdx != 4 {
		t.Error("The found index: ", fIdx, "should be 9")
	}
	if fVal != 14 {
		t.Error("The found val: ", fVal, "should be 14")
	}
}

func TestPartialListFist(t *testing.T) {
	data := []int{80, 36, 89, 14, 87, 66, 86, 49, 56, 26}
	fIdx, fVal := findMin(data, 3)

	if fIdx != 3 {
		t.Error("The found index: ", fIdx, "should be 9")
	}
	if fVal != 14 {
		t.Error("The found val: ", fVal, "should be 14")
	}
}

func TestPartialListLast(t *testing.T) {
	data := []int{80, 36, 89, 26, 87, 66, 86, 49, 56, 14}
	fIdx, fVal := findMin(data, 3)

	if fIdx != 9 {
		t.Error("The found index: ", fIdx, "should be 9")
	}
	if fVal != 14 {
		t.Error("The found val: ", fVal, "should be 14")
	}
}

func TestPartialListMiddle(t *testing.T) {
	data := []int{80, 36, 89, 26, 87, 66, 14, 49, 56, 86}
	fIdx, fVal := findMin(data, 3)

	if fIdx != 6 {
		t.Error("The found index: ", fIdx, "should be 9")
	}
	if fVal != 14 {
		t.Error("The found val: ", fVal, "should be 14")
	}
}
