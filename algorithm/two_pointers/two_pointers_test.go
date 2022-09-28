package two_pointers

import "testing"

func Test_reverseWords(t *testing.T) {

	str := "Let's take LeetCode contest"
	expected := "s'teL ekat edoCteeL tsetnoc"
	result := reverseWords(str)

	if result != expected {
		t.Fatalf("want [%s] but got [%s]", result, expected)
	}

}

func Test_reverseStr(t *testing.T) {

	str := "abcdefg"
	expected := "bacdfeg"
	result := reverseStr(str, 2)

	if result != expected {
		t.Fatalf("want [%s] but got [%s]", expected, result)
	}

	str = "abcd"
	expected = "bacd"
	result = reverseStr(str, 2)

	if result != expected {
		t.Fatalf("want [%s] but got [%s]", result, expected)
	}

}

func Test_reversePrefix(t *testing.T) {

	str := "abcdefd"
	expected := "dcbaefd"
	result := reversePrefix(str, 'd')

	if result != expected {
		t.Fatalf("want [%s] but got [%s]", expected, result)
	}
}
