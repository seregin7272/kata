package sort

import (
	"reflect"
	"sort"
	"testing"
)

func Test_selectionSort(t *testing.T) {
	nums := []int{9, 8, 5, 3, 1, 7, 0, 6, 2, 4}
	expected := []int{9, 8, 5, 3, 1, 7, 0, 6, 2, 4}
	sort.Ints(expected)

	result := selectionSort(nums)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func Test_insertionSort(t *testing.T) {
	nums := []int{9, 8, 5, 3, 1, 7, 0, 6, 2, 4}
	expected := []int{9, 8, 5, 3, 1, 7, 0, 6, 2, 4}
	sort.Ints(expected)

	result := insertionSort(nums)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func Test_bubbleSort(t *testing.T) {
	nums := []int{8, 5, 0, 3, 1, 7, 9, 6, 2, 4}
	expected := []int{8, 5, 0, 3, 1, 7, 9, 6, 2, 4}
	sort.Ints(expected)

	bubbleSort(nums)

	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("expected %v but got %v", expected, nums)
	}
}

func Test_mergeSlice(t *testing.T) {

	a := []int{1, 2, 5, 7, 9}
	b := []int{0, 3, 4, 6, 8}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := mergeSlice(a, b)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v but got %v", expected, result)
	}

}

func Test_mergeSort(t *testing.T) {
	nums := []int{8, 5, 0, 3, 1, 7, 9, 6, 2, 4}
	expected := []int{8, 5, 0, 3, 1, 7, 9, 6, 2, 4}
	sort.Ints(expected)

	result := mergeSort(nums)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v but got %v", expected, result)
	}

}

func Test_Quick(t *testing.T) {
	nums := []int{8, 5, 0, 3, 1, 7, 9, 6, 2, 4}
	expected := []int{8, 5, 0, 3, 1, 7, 9, 6, 2, 4}
	sort.Ints(expected)

	quickSort(nums, 0, len(nums))

	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("expected %v but got %v", expected, nums)
	}

}

func Test_CountSort(t *testing.T) {
	nums := []int{8, 5, 0, 3, 1, 7, 9, 6, 2, 4}
	expected := []int{8, 5, 0, 3, 1, 7, 9, 6, 2, 4}
	sort.Ints(expected)

	countSort(nums)

	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("expected %v but got %v", expected, nums)
	}

}
