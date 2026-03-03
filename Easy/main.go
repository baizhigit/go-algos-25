package main

import (
	"fmt"
	"math"
)

func firstNPrimes(n int) []int {
	primes := []int{}
	for num := 2; len(primes) < n; num++ {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}
	return primes
}

// IsPrime проверяет, является ли число простым
// Временная сложность: O(√n)
// Пространственная сложность: O(1)
func isPrime(n int) bool {
	// 1. Отсекаем числа меньше 2
	if n < 2 {
		return false
	}

	// 2. Быстрая проверка для самых маленьких простых чисел
	if n == 2 || n == 3 {
		return true
	}

	// 3. Отсекаем чётные и кратные 3
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// 4. Перебираем делители вида 6k ± 1 до √n
	// Проверка i*i <= n гарантирует, что мы проверяем делители только до √n, потому что если у числа есть делитель больше √n,
	// то второй делитель будет меньше √n, и мы его уже найдём. Умножение i*i работает быстрее и точнее, чем использование sqrt(n)
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

// 1. Reverse array
func reverseArray(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

// 2. Reverse string
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 3. Чётное или нечётное число
func isEven(n int) bool { return n%2 == 0 }

// 4. Max in array
func maxInArray(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

// 5. Min in array
func minInArray(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

// 6. Sum 1..N
func sumToN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// Сумма чисел от 1 до N
func sumToNFast(n int) int {
	return n * (n + 1) / 2
}

// 7. Sum array
func sumArray(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 8. Palindrome string
func isPalindrome(s string) bool {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

// 9. Count char
func countChar(s string, target rune) int {
	count := 0
	for _, r := range s {
		if r == target {
			count++
		}
	}
	// if target is byte
	// for i := 0; i < len(s); i++ {
	// 	if s[i] == target {
	// 		count++
	// 	}
	// }
	return count
}

// 10. Power of two
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// 11. Second largest
func secondLargest(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	max1, max2 := math.MinInt, math.MinInt
	for _, v := range nums {
		if v > max1 {
			max2 = max1
			max1 = v
		} else if v > max2 && v != max1 {
			max2 = v
		}
	}
	if max2 == math.MinInt {
		return -1
	}
	return max2
}

// 12. Найти все уникальные элементы массива
func unique(nums []int) []int {
	seen := map[int]bool{}
	res := []int{}
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			res = append(res, v)
		}
	}
	return res
}

// 13. Intersection of arrays
func intersection(a, b []int) []int {
	m := map[int]bool{}
	res := []int{}
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if m[v] {
			res = append(res, v)
			delete(m, v)
		}
	}
	return res
}

// 14. Merge sorted arrays - Объединение двух отсортированных массивов
func mergeSorted(a, b []int) []int {
	i, j := 0, 0
	res := []int{}
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res
}

// 15. Linear search
func linearSearch(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}

// 16. Binary search
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

// 19. Fibonacci
func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 20. Factorial
func factorial(n int) int {
	out := 1
	for i := 2; i <= n; i++ {
		out *= i
	}
	return out
}

// 21. Two Sum
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return nil
}

func main() {

	fmt.Println("firstNPrimes: ", firstNPrimes(15))
	fmt.Println("isEven: ", isEven(7))
	fmt.Println("reverseArray: ", reverseArray([]int{1, 2, 3, 4, 5, 8}))
	fmt.Println("reverseString: ", reverseString("kamolino"))
	fmt.Println("maxInArray: ", maxInArray([]int{1, 12, 3, 4, 5, 8}))
	fmt.Println("minInArray: ", minInArray([]int{18, 12, 3, 4, 5, 8}))
	fmt.Println("sumToNFast: ", sumToNFast(7))
	fmt.Println("sumToN: ", sumToN(7))
	fmt.Println("sumArray: ", sumArray([]int{18, 12, 3, 4, 5, 8}))
	fmt.Println("isPalindrome: ", isPalindrome("kamolino"))
	fmt.Println("isPalindrome: ", isPalindrome("rotator"))
	fmt.Println("countChar: ", countChar("rotator", 't'))
	fmt.Println("isPowerOfTwo: ", isPowerOfTwo(9))
	fmt.Println("secondLargest: ", secondLargest([]int{18, 12, 3, 4, 5, 8}))
	fmt.Println("secondLargest: ", secondLargest([]int{2, 2}))
	fmt.Println("unique: ", unique([]int{2, 3, 3, 4, 6, 6, 8, 5, 8}))
	fmt.Println("intersection: ", intersection([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}))
	fmt.Println("mergeSorted: ", mergeSorted([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}))
	fmt.Println("linearSearch: ", linearSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4))
	fmt.Println("binarySearch: ", binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2))
	fmt.Println("fib: ", fib(6), fib(7), fib(8), fib(9))
	fmt.Println("factorial: ", factorial(2), factorial(3), factorial(4), factorial(5))
	fmt.Println("twoSum: ", twoSum([]int{2, 7, 11, 15}, 9))

}
