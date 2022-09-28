package bit

import (
	"reflect"
	"testing"
)

func Test_Subtest(t *testing.T) {
	expected := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
	nums := []int{1, 2, 3}
	result := subsets(nums)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("result not match got %v but want %v", result, expected)
	}
}

func Test_CountConsistentStrings(t *testing.T) {

	word1 := []string{"ad", "bd", "aaab", "baa", "badab"}
	allowed := "ab"
	result := countConsistentStrings(allowed, word1)
	expected := 2

	if expected != result {
		t.Fatalf("result not match got %v but want %v", result, expected)
	}

	word2 := []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
	allowed = "cad"
	result = countConsistentStrings(allowed, word2)
	expected = 4

	if expected != result {
		t.Fatalf("result not match got %v but want %v", result, expected)
	}

}

func Test_subsetXORSum(t *testing.T) {

	nums := []int{5, 1, 6}
	expected := 28
	result := subsetXORSum(nums)

	if result != expected {
		t.Fatalf("result not match got %v but want %v", result, expected)
	}

}

func Test_Bits(t *testing.T) {

	bits()

}
