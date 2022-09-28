package main

import (
	"fmt"
)

type Stack struct {
	items []string
}

func (s *Stack) push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) top() string {
	if s.isEmpty() {
		return ""
	}

	return s.items[len(s.items)-1]
}

func (s *Stack) pop() error {
	if s.isEmpty() {
		return fmt.Errorf("emty stack")
	}
	s.items = s.items[:len(s.items)-1]

	return nil
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func isBalanced(chars string) bool {
	isOpen := func(s string) bool {
		return s == "[" || s == "{" || s == "("
	}

	isPair := func(s, open string) bool {
		return open == "[" && s == "]" ||
			open == "{" && s == "}" || open == "(" && s == ")"
	}

	stack := &Stack{}
	for _, char := range chars {
		strChar := string(char)
		if isOpen(strChar) {
			stack.push(strChar)
		} else {

			if isPair(strChar, stack.top()) {
				stack.pop()
			} else {
				return false
			}

		}
	}

	return stack.isEmpty()
}

func main() {
	s1 := "[()]"
	s2 := "[()"
	s3 := "{()}"
	s4 := "[(])"
	s5 := "()([](){}"

	fmt.Println(isBalanced(s1)) // true
	fmt.Println(isBalanced(s2)) // false
	fmt.Println(isBalanced(s3)) // true
	fmt.Println(isBalanced(s4)) // false
	fmt.Println(isBalanced(s5)) // false
}
