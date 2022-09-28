package combinatorics

import "sort"

// https://leetcode.com/problems/permutations/
// Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
//Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

func permute(nums []int) [][]int {
	var result [][]int

	if len(nums) == 1 {
		result = append(result, nums)
		return result
	}

	for i, val := range nums {
		var nums2 []int
		nums2 = append(nums2, nums[:i]...)
		nums2 = append(nums2, nums[i+1:]...)
		p := permute(nums2)
		for _, s := range p {
			tmp := []int{val}
			tmp = append(tmp, s...)
			result = append(result, tmp)
		}

	}

	return result
}

// https://leetcode.com/problems/generate-parentheses/
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]

// Input: n = 2
// Output: ["((()))","(()())","(())()","()(())","()()()"]
func generateParenthesis(n int) []string {

	res := make([]string, 0, n)
	genBalance(0, &res, 0, "", n)

	return res
}

func genBalance(balance int, res *[]string, pos int, str string, n int) {
	if n*2 == pos {
		*res = append(*res, str)
	} else {

		if balance > 0 {
			genBalance(balance-1, res, pos+1, str+")", n)
		}

		if balance+1 <= n*2-(pos+1) {
			genBalance(balance+1, res, pos+1, str+"(", n)
		}
	}
}

// https://leetcode.com/problems/build-array-from-permutation/

//Example 1:
//
//Input: nums = [0,2,1,5,3,4]
//Output: [0,1,2,4,5,3]
//Explanation: The array ans is built as follows:
//ans = [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]]
//= [nums[0], nums[2], nums[1], nums[5], nums[3], nums[4]]
//= [0,1,2,4,5,3]

//Example 2:
//
//Input: nums = [5,0,1,2,3,4]
//Output: [4,5,0,1,2,3]
//Explanation: The array ans is built as follows:
//ans = [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]]
//= [nums[5], nums[0], nums[1], nums[2], nums[3], nums[4]]
//= [4,5,0,1,2,3]
func buildArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		ans[i] = nums[val]
	}

	return ans
}

// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/

//Example 1:
//
//Input: seats = [3,1,5], students = [2,7,4]
//Output: 4
//Explanation: The students are moved as follows:
//- The first student is moved from from position 2 to position 1 using 1 move.
//- The second student is moved from from position 7 to position 5 using 2 moves.
//- The third student is moved from from position 4 to position 3 using 1 move.
//In total, 1 + 2 + 1 = 4 moves were used.

//Example 2:
//
//Input: seats = [4,1,5,9], students = [1,3,2,6]
//Output: 7
//Explanation: The students are moved as follows:
//- The first student is not moved.
//- The second student is moved from from position 3 to position 4 using 1 move.
//- The third student is moved from from position 2 to position 5 using 3 moves.
//- The fourth student is moved from from position 6 to position 9 using 3 moves.
//In total, 0 + 1 + 3 + 3 = 7 moves were used.

//Example 3:
//
//Input: seats = [2,2,6,6], students = [1,3,2,6]
//Output: 4
//Explanation: Note that there are two seats at position 2 and two seats at position 6.
//The students are moved as follows:
//- The first student is moved from from position 1 to position 2 using 1 move.
//- The second student is moved from from position 3 to position 6 using 3 moves.
//- The third student is not moved.
//- The fourth student is not moved.
//In total, 1 + 3 + 0 + 0 = 4 moves were used.

func minMovesToSeat(seats []int, students []int) int {
	ans := 0

	sort.Ints(seats)
	sort.Ints(students)

	for i := 0; i < len(seats); i++ {
		if seats[i] > students[i] {
			ans += seats[i] - students[i]
		} else {
			ans += students[i] - seats[i]
		}
	}

	return ans

}
