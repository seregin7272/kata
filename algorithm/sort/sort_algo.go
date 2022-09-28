package sort

import (
	"math/rand"
	"time"
)

// Находим самый наименьший и меняем с первым
// O(n^2) временная сложность
// 0(1) памяти
func selectionSort(nums []int) []int {
	// перебираем массив
	// n + (n-1) + (n-2) ... + 1 = n * (n-1) / 2 = 0(n^2)
	for i, _ := range nums {
		minIdx := i
		// ищем индекс минимального элемента в массиве
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		// меняеи местами первый и минимальный элемент
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}

// Сравниваем элмент на i и (i-1) если i < (i -1) то менеям им местами,
// далее идем по уже ранее отсортировнным элементам в начале массива
// и сортируем этот префикс с учетом нового элемента
// O(n^2) время
// 0(1) памяти
func insertionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}

	return nums
}

// на каждой итерации перебираем все элементы массива,
// и если соседние образуют инверсию, то меняем их местами.
// Итерации делаем до тех пор, пока есть инверсии.
// O(n^2) время
// 0(1) памяти
func bubbleSort(nums []int) {
	isSorted := true
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				isSorted = false
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
		if isSorted {
			return
		}
	}
}

// делит массив на две примерно равных половины и делает рекурсивные запуски для обеих частей,
// после чего сливает отсортированные половины
// в один отсортированный массив и возвращает в качестве результата
// O(n*log(n))
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	c := len(nums) / 2
	a := nums[:c]
	b := nums[c:]

	return mergeSlice(mergeSort(a), mergeSort(b))
}

func mergeSlice(a, b []int) []int {
	var result []int
	l, r := 0, 0

	for {
		if a[l] < b[r] {
			result = append(result, a[l])
			l++
		} else {
			result = append(result, b[r])
			r++
		}

		if l == len(a) {
			result = append(result, b[r:]...)
			break
		}

		if r == len(b) {
			result = append(result, a[l:]...)
			break
		}
	}

	return result
}

// До рекурсивных запусков партиционируем массив так,
// чтобы выбрать опорный элемент и перенести все меньшие опорного влево, остальные — вправо.
// После рекурсивных запусков делать ничего не нужно.
func quickSort(nums []int, l int, r int) {
	if l == r {
		return
	}

	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(r-l) + l
	nums[rand], nums[l] = nums[l], nums[rand]

	s1SectionIdx := l
	s2SectionIdx := l

	// партиционируем массив слева то что меньше p, а справа больше
	p := l

	for k := l + 1; k < r; k++ {
		if nums[k] >= nums[p] {
			s2SectionIdx++
		} else {
			s1SectionIdx++
			nums[s1SectionIdx], nums[k] = nums[k], nums[s1SectionIdx]
			s2SectionIdx++
		}
	}

	nums[p], nums[s1SectionIdx] = nums[s1SectionIdx], nums[p]
	p = s1SectionIdx

	quickSort(nums, l, p)
	quickSort(nums, p+1, r)
}

func countSort(nums []int) {
	max := nums[0]

	for _, val := range nums {
		if val > max {
			max = val
		}
	}

	bucket := make([]int, max+1)

	for _, val := range nums {
		bucket[val]++
	}

	nums = nums[:0]

	for j, val := range bucket {
		count := bucket[val]
		for i := 0; i < count; i++ {
			nums = append(nums, j)
		}

	}
}
