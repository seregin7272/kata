package sort

import (
	"reflect"
	"testing"
)

func Test_minimumSum(t *testing.T) {

	num := 2932
	expected := 52
	result := minimumSum(num)

	if expected != result {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func Test_smallerNumbersThanCurrent(t *testing.T) {

	num := []int{8, 1, 2, 2, 3}
	expected := []int{4, 0, 1, 1, 3}
	result := smallerNumbersThanCurrent(num)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v but got %v", expected, result)
	}
}
