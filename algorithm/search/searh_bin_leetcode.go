package search

import "fmt"

func search(nums []int, target int) int {
	l := 0
	r := len(nums)

	fmt.Printf("BEGIN target = %d  l = %d  r = %d \n\n\n", target, l, r)

	for l != r-1 {
		mid := (l + r) / 2
		fmt.Printf("FOR1 mid = %d  val =  %d \n", mid, nums[mid])

		if nums[mid] > target {
			r = mid // двигаем правую границу
		} else {
			l = mid // двигаем левую границу
		}

		fmt.Printf("FOR2 l = %d  r = %d \n\n\n", l, r)
	}

	if nums[l] == target {
		return l
	}

	return -1
}

func search2(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	fmt.Printf("BEGIN target = %d  l = %d  r = %d \n\n\n", target, l, r)

	for l <= r {
		mid := (l + r) / 2

		fmt.Printf("FOR1 mid = %d  val =  %d \n", mid, nums[mid])

		if nums[mid] == target {
			l = mid
			fmt.Printf("return mid = %d  val =  %d \n", mid, nums[mid])
			return l
		}

		if nums[mid] > target {
			r = mid - 1 // двигаем правую границу
		} else {
			l = mid + 1 // двигаем левую границу
		}

		fmt.Printf("FOR2 l = %d  r = %d \n\n\n", l, r)
	}

	return -1
}

//https://leetcode.com/problems/find-target-indices-after-sorting-array/
//You are given a 0-indexed integer array nums and a target element target.
//
//A target index is an index i such that nums[i] == target.
//
//Return a list of the target indices of nums after sorting nums in non-decreasing order. If there are no target indices, return an empty list. The returned list must be sorted in increasing order.
//
//
//
//Example 1:
//
//Input: nums = [1,2,5,2,3], target = 2
//Output: [1,2]
//Explanation: After sorting, nums is [1,2,2,3,5].
//The indices where nums[i] == 2 are 1 and 2.
//Example 2:
//
//Input: nums = [1,2,5,2,3], target = 3
//Output: [3]
//Explanation: After sorting, nums is [1,2,2,3,5].
//The index where nums[i] == 3 is 3.
//Example 3:
//
//Input: nums = [1,2,5,2,3], target = 5
//Output: [4]
//Explanation: After sorting, nums is [1,2,2,3,5].
//The index where nums[i] == 5 is 4.

func targetIndices(nums []int, target int) []int {
	var result []int

	// сортируем массив от меньшего к большему
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}

	fmt.Println(nums)
	// через бинарный поиск ищем самый левый индекс

	l := -1
	r := len(nums) - 1

	for l != r-1 {
		mid := (l + r) / 2

		if nums[mid] >= target {
			r = mid
		} else {
			l = mid
		}
	}

	if nums[r] != target {
		return result
	}

	// проверяем последователь элементы правее и если они равны элементы под найденным индексом
	// добавляем их в результат
	result = append(result, r)
	for i := r + 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			break
		}
		result = append(result, i)
	}

	fmt.Println(result)

	return result
}

// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must be unique and you may return the result in any order.
//Example 1:
//
//Input: nums1 = [1,2,2,1], nums2 = [2,2]
//Output: [2]
//Example 2:
//
//Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//Output: [9,4]
//Explanation: [4,9] is also accepted.

func intersection(nums1 []int, nums2 []int) []int {
	var result []int

	m := map[int]struct{}{}

	for i := 0; i < len(nums1); i++ {
		for j := i; j > 0; j-- {
			if nums1[j] < nums1[j-1] {
				nums1[j-1], nums1[j] = nums1[j], nums1[j-1]
			}
		}
	}

	for _, val := range nums2 {
		if _, ok := m[val]; ok {
			continue
		}

		l := 0
		r := len(nums1)

		for l != r-1 {
			mid := (l + r) / 2

			if nums1[mid] > val {
				r = mid
			} else {
				l = mid
			}
		}

		if nums1[l] == val {
			result = append(result, val)
		}

		m[val] = struct{}{}
	}

	return result
}

//Input: n = 5
//Output: 2
//Explanation: Because the 3rd row is incomplete, we return 2.

func arrangeCoins(n int) int {
	var row int

	for n > 0 {

		row++
		n -= row

		if n < 0 {
			row--
		}
	}

	return row
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
//Example 1:
//
//Input: nums = [1,1,2]
//Output: 2, nums = [1,2,_]
//Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
//It does not matter what you leave beyond the returned k (hence they are underscores).
//Example 2:
//
//Input: nums = [0,0,1,1,1,2,2,3,3,4]
//Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
//Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
//It does not matter what you leave beyond the returned k (hence they are underscores).

func removeDuplicates(nums []int) int {
	l := 1

	for r := 1; r < len(nums); r++ {
		if nums[l-1] == nums[r] {
			continue
		}

		nums[l], l = nums[r], l+1

	}

	return l
}

// https://leetcode.com/problems/peak-index-in-a-mountain-array/
//Массив arr гора , если выполняются следующие свойства:
//
//arr.length >= 3
//Существуют некоторые iс 0 < i < arr.length - 1таким, что:
//arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
//Учитывая горный массив arr, вернуть индекс i таким образом, что arr[0] < arr[1] < ... < arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].
//
//Вы должны решить его по O(log(arr.length))временной сложности.
////Example 1:
//
//Input: arr = [0,1,0]
//Output: 1
//Example 2:
//
//Input: arr = [0,2,1,0]
//Output: 1
//Example 3:
//
//Input: arr = [0,10,5,2]
//Output: 1
func peakIndexInMountainArray(arr []int) int {
	var ans int

	return ans
}

// https://leetcode.com/problems/find-peak-element/
// Example 1:
//
//Input: nums = [1,2,3,1]
//Output: 2
//Explanation: 3 is a peak element and your function should return the index number 2.
//Example 2:
//
//Input: nums = [1,2,1,3,5,6,4]
//Output: 5
//Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.
func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}

	}

	return l
}

// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
//Example 1:
//
//Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
//Output: 8
//Explanation: There are 8 negatives number in the matrix.

//Example 2:
//
//Input: grid = [[3,2],[1,0]]
//Output: 0
func countNegatives(grid [][]int) int {
	var ans int

	for _, item := range grid {
		l := -1
		r := len(item) - 1

		for l < r-1 {
			mid := (l + r) / 2
			if item[mid] < 0 {
				r = mid
			} else {
				l = mid
			}
		}

		if item[r] < 0 {
			ans += len(item) - r
		}

	}

	return ans
}
