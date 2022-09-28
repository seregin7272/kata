package bit

import "fmt"

//
//https://leetcode.com/problems/subsets/
//Given an integer array nums of unique elements, return all possible subsets (the power set).
//The solution set must not contain duplicate subsets. Return the solution in any order.
//Example 1:
//
//Input: nums = [1,2,3]
//Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

func subsets(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	// перебор всех вариантов масок
	for mask := 0; mask < (1 << n); mask++ {
		var t []int
		// перебор всех битов для маски
		for i := 0; i < n; i++ {
			// проверка есть ли на i месте в этой маске 1, если не 0 значит на этом месте есть значение из множеества
			if (mask & (1 << i)) != 0 {
				t = append(t, nums[i])
			}
		}

		result = append(result, t)
	}

	return result
}

//https://leetcode.com/problems/count-the-number-of-consistent-strings/
//You are given a string allowed consisting of distinct characters and an array of strings words. A string is consistent if all characters in the string appear in the string allowed.
//
//Return the number of consistent strings in the array words.
//
//
//
//Example 1:
//
//Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
//Output: 2
//Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.
//Example 2:
//
//Input: allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
//Output: 7
//Explanation: All strings are consistent.
//Example 3:
//
//Input: allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
//Output: 4
//Explanation: Strings "cc", "acd", "ac", and "d" are consistent.
func countConsistentStrings(allowed string, words []string) int {
	maskAllowed := 0
	cnt := 0

	// дизъюнкции (или)
	// добавить элементы одного множества в другое
	for _, letter := range allowed {
		maskAllowed = maskAllowed | 1<<(letter-'a')
	}

	for _, word := range words {
		cnt++
		for _, letter := range word {
			n := 1 << (letter - 'a')
			// конъюнкции (и)
			// входит ли буква в подмножесво если нет то и все слово не является подмножеством allowed
			if (maskAllowed & n) == 0 {
				cnt--
				break
			}
		}
	}

	return cnt
}

// https://leetcode.com/problems/sum-of-all-subset-xor-totals/
//Example 1:
//
//Input: nums = [1,3]
//Output: 6
//Explanation: The 4 subsets of [1,3] are:
//- The empty subset has an XOR total of 0.
//- [1] has an XOR total of 1.
//- [3] has an XOR total of 3.
//- [1,3] has an XOR total of 1 XOR 3 = 2.
//0 + 1 + 3 + 2 = 6

//Example 2:
//
//Input: nums = [5,1,6]
//Output: 28
//Explanation: The 8 subsets of [5,1,6] are:
//- The empty subset has an XOR total of 0.
//- [5] has an XOR total of 5.
//- [1] has an XOR total of 1.
//- [6] has an XOR total of 6.
//- [5,1] has an XOR total of 5 XOR 1 = 4.
//- [5,6] has an XOR total of 5 XOR 6 = 3.
//- [1,6] has an XOR total of 1 XOR 6 = 7.
//- [5,1,6] has an XOR total of 5 XOR 1 XOR 6 = 2.
//0 + 5 + 1 + 6 + 4 + 3 + 7 + 2 = 28

//Example 3:
//
//Input: nums = [3,4,5,6,7,8]
//Output: 480
//Explanation: The sum of all XOR totals for every subset is 480.

func subsetXORSum(nums []int) int {
	n := len(nums)
	var ans int

	for i := 0; i < (1 << n); i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				if sum <= 0 {
					sum = nums[j]
				} else {
					sum = sum ^ nums[j]
				}

			}
		}

		ans += sum
	}

	return ans
}

// https://leetcode.com/problems/decode-xored-array
//Input: encoded = [1,2,3], first = 1
//Output: [1,0,2,1]
//Explanation: If arr = [1,0,2,1], then first = 1 and encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [1,2,3]
// Так как encoded[i] = arr[i] XOR arr[i+1], тогда arr[i+1] = encoded[i] XOR arr[i].
func decode(encoded []int, first int) []int {
	ans := make([]int, len(encoded)+1)
	ans[0] = first

	for i, val := range encoded {
		ans[i+1] = val ^ ans[i]
	}

	return ans
}

func bits() {

	//sets := []int{1, 0, 0, 0, 1, 0, 0, 1} // 137
	setN := 137

	findN3 := 8  // 1000 // индекс 3
	findN5 := 32 // 100000 // индекс 5

	// конъюнкции (и)
	if (setN & findN3) != 0 {
		fmt.Println("индекс 3 занят")
	}

	if (setN & (1 << 3)) != 0 {
		fmt.Println("индекс 3 занят")
	}

	if (setN & findN5) == 0 {
		fmt.Println("индекс 5 не занят")
	}

	// дизъюнкции (или)
	// добавить элементы одного множества в другое
	setN = setN | (1 << 5) // setN = 169 / 10101001

	if (setN & findN5) != 0 {
		fmt.Println("индекс 5 занят")
	}

	//  XOR (Исключающее «или»)
	// Удаляет элементы одного множества из другого
	setN2 := setN ^ findN3 // setN2 = 161 / 10100001

	fmt.Println(setN2)

	fmt.Println(10 << 1) // 20
	fmt.Println(10 << 2) // 40
	fmt.Println(10 << 3) // 60

	fmt.Println(2 ^ 5 ^ 6) // 1
	fmt.Println("     [    ")
	fmt.Println(1 ^ 0) // 1
	fmt.Println(0 ^ 2) // 2
	fmt.Println(2 ^ 1) // 3
	fmt.Println("    ]     ")
	fmt.Println(1 ^ 1)
	fmt.Println(0 ^ 2)
	fmt.Println(2 ^ 3)

}
