package search

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_search(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expected := 4
	result := search(nums, target)

	if result != expected {
		t.Fatalf("expected want %d, got %d", expected, target)
	}

	target = 2
	expected = -1
	result = search(nums, target)

	if result != expected {
		t.Fatalf("expected want %d, got %d", expected, target)
	}
}

func Test_search2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expected := 4
	result := search2(nums, target)

	if result != expected {
		t.Fatalf("expected want %d, got %d", expected, target)
	}

	target = 2
	expected = -1
	result = search2(nums, target)

	if result != expected {
		t.Fatalf("expected want %d, got %d", expected, target)
	}

	nums = []int{5}
	target = 5
	expected = 0
	result = search2(nums, target)

	if result != expected {
		t.Fatalf("expected want %d, got %d", expected, target)
	}
}

func Test_targetIndices(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{1, 2, 5, 2, 3},
			target:   2,
			expected: []int{1, 2},
		},
		{
			nums:     []int{48, 90, 9, 21, 31, 35, 19, 69, 29, 52, 100, 54, 21, 86, 6, 45, 42, 5, 62, 77, 15, 38},
			target:   6,
			expected: []int{1},
		},
	}

	for i, val := range cases {
		result := targetIndices(val.nums, val.target)
		if !reflect.DeepEqual(val.expected, result) {
			t.Fatalf("case %d wanted %v, but got %v", i, val.expected, result)
		}
	}

}

func Test_intersection(t *testing.T) {

	cases := []struct {
		nums1    []int
		nums2    []int
		target   int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2},
		},
		{
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{9, 4},
		},
	}

	for i, val := range cases {
		result := intersection(val.nums1, val.nums2)
		if !reflect.DeepEqual(val.expected, result) {
			t.Fatalf("case %d wanted %v, but got %v", i, val.expected, result)
		}
	}

}

func Test_removeDuplicates(t *testing.T) {
	cases := []struct {
		nums          []int
		expectedCount int
		expected      []int
	}{
		{
			nums:          []int{1, 1, 2},
			expectedCount: 2,
			expected:      []int{1, 2, 1},
		},
		{
			nums:          []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedCount: 5,
			expected:      []int{0, 1, 2, 3, 4, 1, 1, 1, 1, 1},
		},
	}

	for i, val := range cases {
		result := removeDuplicates(val.nums)
		fmt.Println(val.nums)
		if result != val.expectedCount {
			t.Fatalf("case %d wanted %v, but got %v", i, val.expected, result)
		}
	}

}
