package sort

import "sort"

//https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
//
//Example 1:
//
//Input: num = 2932
//Output: 52
//Explanation: Some possible pairs [new1, new2] are [29, 23], [223, 9], etc.
//The minimum sum can be obtained by the pair [29, 23]: 29 + 23 = 52.
//Example 2:
//
//Input: num = 4009
//Output: 13
//Explanation: Some possible pairs [new1, new2] are [0, 49], [490, 0], etc.
//The minimum sum can be obtained by the pair [4, 9]: 4 + 9 = 13.

func minimumSum(num int) int {
	var nums []int

	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}

	sort.Ints(nums)

	return nums[0]*10 + nums[2] + nums[1]*10 + nums[3]
}

// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
//Input: nums = [1,2,2,3,8]
//Output: [4,0,1,1,3]
//Explanation:
//For nums[0]=8 there exist four smaller numbers than it (1, 2, 2 and 3).
//For nums[1]=1 does not exist any smaller number than it.
//For nums[2]=2 there exist one smaller number than it (1).
//For nums[3]=2 there exist one smaller number than it (1).
//For nums[4]=3 there exist three smaller numbers than it (1, 2 and 2).
func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))

	max := nums[0]
	for _, val := range nums {
		if val > max {
			max = val
		}
	}

	bucket := make([]int, max+1)
	// [8 1 2 2 3]
	// [0 1 2 3 4 5 6 7 8]
	// [0 1 2 1 0 0 0 0 1]
	for _, val := range nums {
		bucket[val]++
	}

	// [0 1 3 4 4 4 4 4 5]
	for i := 1; i < len(bucket); i++ {
		bucket[i] += bucket[i-1]
	}

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if val == 0 {
			result[i] = 0
		} else {
			result[i] = bucket[val-1]
		}
	}

	return result
}
