package two_pointers

// https://leetcode.com/problems/reverse-words-in-a-string-iii/
// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
// Example 2:
//
// Input: s = "God Ding"
// Output: "doG gniD"
func reverseWords(s string) string {
	spaceIdx := -1

	ss := []byte(s)

	for strIdx := 0; strIdx <= len(ss); strIdx++ {
		if strIdx == len(ss) || ss[strIdx] == ' ' {
			startIdx := spaceIdx + 1
			endIdx := strIdx - 1
			for startIdx < endIdx {
				tmp := ss[startIdx]
				ss[startIdx] = ss[endIdx]
				ss[endIdx] = tmp
				startIdx++
				endIdx--
			}
			spaceIdx = strIdx
		}
	}

	return string(ss)
}

// https://leetcode.com/problems/reverse-string-ii/
//Example 1:
//
//Input: s = "abcdefg", k = 2
//Output: "bacdfeg"

//Example 2:
//
//Input: s = "abcd", k = 2
// Output: "bacd"
func reverseStr(s string, k int) string {
	ss := []byte(s)

	for i := 0; i < len(ss); i += 2 * k {
		l := i
		r := i + k - 1

		if (len(ss) - 1) < r {
			r = len(ss) - 1
		}

		for l < r {
			t := ss[l]
			ss[l] = ss[r]
			ss[r] = t
			l++
			r--
		}

	}

	return string(ss)
}

// https://leetcode.com/problems/reverse-prefix-of-word/
//Example 1:
//
//Input: word = "abcdefd", ch = "d"
//Output: "dcbaefd"
//Explanation: The first occurrence of "d" is at index 3.
//Reverse the part of word from 0 to 3 (inclusive), the resulting string is "dcbaefd".
//Example 2:
//
//Input: word = "xyxzxe", ch = "z"
//Output: "zxyxxe"
//Explanation: The first and only occurrence of "z" is at index 3.
//Reverse the part of word from 0 to 3 (inclusive), the resulting string is "zxyxxe".
//Example 3:
//
//Input: word = "abcd", ch = "z"
//Output: "abcd"
//Explanation: "z" does not exist in word.
//You should not do any reverse operation, the resulting string is "abcd".
func reversePrefix(word string, ch byte) string {
	ss := []byte(word)

	for i := 0; i < len(ss); i++ {
		if ss[i] == ch {
			l := 0
			r := i

			for l < r {
				t := ss[l]
				ss[l] = ss[r]
				ss[r] = t
				r--
				l++
			}

			break
		}

	}

	return string(ss)
}
