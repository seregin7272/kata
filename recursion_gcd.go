package main

import (
	"fmt"
	"log"
)

func main() {

	r2 := pow(2, 8)

	log.Print(r2)

}

func hanoi(n int, from int, to int) {
	if n == 1 {
		fmt.Printf("from %d to %d", from, to)
		fmt.Println()
		return
	}
	// подсчет номер промежуточной башни
	mid := 6 - from - to

	hanoi(n-1, from, mid)
	fmt.Printf("from %d to %d", from, to)
	fmt.Println()
	hanoi(n-1, mid, to)
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

//  НОД вычитанием O(min(a,b))
// много опрераций при 10 в степени 9
func gcd1(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	if a == b {
		return a
	}

	if a > b {
		a = a - b
	} else {
		b = b - a
	}

	return gcd1(a, b)
}

// НОД через остаток от деления O(log(min(a,b))
func gcd2(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	if a > b {
		a = a % b
	} else {
		b = b % a
	}

	return gcd2(a, b)
}

// НОД для массива O(N * log(N))
func findGCD(nums []int, n int) int {
	if n == len(nums)-1 {
		return nums[n]
	}

	a := nums[n]
	b := findGCD(nums, n+1)

	return gcd2(a, b)
}

// НОК 2 чисел
func lcm(a, b int) int {
	return a * b / gcd1(a, b)
}

// НОК для массива O(n * log(min(a, b)))
func findLCM(nums []int, n int) int {
	if n == len(nums)-1 {
		return nums[n]
	}

	a := nums[n]
	b := findLCM(nums, n+1)

	return lcm(a, b)
}

// Быстрое возведение в степень O (Log n)
func pow(a int, b int) int {
	if b == 0 {
		return 1
	}

	x := pow(a, b/2)

	if b%2 == 0 {
		return x * x
	}

	return x * x * a
}

func mypow(a, b int) int {
	if b == 0 {
		return 1
	}

	if b%2 == 0 {
		return mypow(a, b/2) * mypow(a, b/2)
	} else {
		return mypow(a, b/2) * mypow(a, b/2) * a
	}
}
