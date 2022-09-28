package combinatorics

import (
	"reflect"
	"testing"
)

type caseTest struct {
	expected []string
	n        int
}

func Test_Permute(t *testing.T) {
	expected := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	result := permute([]int{1, 2, 3})

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result not match\nGot : %#v\nWant: %#v", result, expected)
	}
}

func Test_GenerateParenthesis(t *testing.T) {

	cases := []caseTest{
		{
			expected: []string{"()"},
			n:        1,
		},
		{
			expected: []string{"()()", "(())"},
			n:        2,
		},
		{
			expected: []string{"()()()", "()(())", "(())()", "(()())", "((()))"},
			n:        3,
		},
	}

	for _, caseItem := range cases {
		result := generateParenthesis(caseItem.n)

		if !reflect.DeepEqual(result, caseItem.expected) {
			t.Fatalf("n = %d, result not match\nGot : %#v\nWant: %#v", caseItem.n, result, caseItem.expected)
		}
	}
}
