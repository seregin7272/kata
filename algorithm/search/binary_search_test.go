package search

import (
	"testing"
)

func Test_binarysearch1(t *testing.T) {

	nums := []int{10, 20, 35, 40, 54}
	n := 5
	y := 40

	// искомый элемент входит в массив
	result := binsearch1(nums, n, y)
	expected := 40

	if result != expected {
		t.Errorf("execpted %d but got %d", expected, result)
	}

	// искомый элемент больше самого правого в массиве
	y = 100
	result = binsearch1(nums, n, y)
	expected = 54

	if result != expected {
		t.Errorf("execpted %d but got %d", expected, result)
	}

	// искомый элемент меньше самого левого в массиве
	y = 9
	result = binsearch1(nums, n, y)
	expected = 0

	if result != expected {
		t.Errorf("execpted %d but got %d", expected, result)
	}

}

func Test_right_left_searchbin(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 5, 5, 7, 42}
	n := len(nums)
	y := 5

	expected := 6
	result := rightbin(nums, n, y)
	if result != expected {
		t.Errorf("execpted %d but got %d", expected, result)
	}

	expected = 4
	result = leftbin(nums, n, y)
	if result != expected {
		t.Errorf("execpted %d but got %d", expected, result)
	}
}

func Test_countEqualItem(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 5, 5, 7, 42}
	n := len(nums)
	y := 5

	expected := 3
	result := countEqualItem(nums, n, y)
	if result != expected {
		t.Errorf("execpted %d but got %d", expected, result)
	}
}

func Test_findPeakElement(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := 2
	result := findPeakElement(nums)

	if result != expected {
		t.Errorf("execpted %d but got %d", expected, result)
	}

	nums = []int{1, 2, 1, 3, 5, 6, 4}
	expected = 5

	result = findPeakElement(nums)

	if result != expected {
		t.Errorf("execpted %d but got %d", expected, result)
	}

}
