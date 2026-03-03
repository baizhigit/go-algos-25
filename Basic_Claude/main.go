package main

import (
	"fmt"
	"math"
	"strings"
)

// ============================================================================
// РАБОТА СО СТРОКАМИ
// ============================================================================

// IsPalindrome проверяет, является ли строка палиндромом
// Временная сложность: O(n), где n - длина строки
// Пространственная сложность: O(1)
func IsPalindrome(s string) bool {
	// Убираем пробелы и приводим к нижнему регистру
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	left, right := 0, len(s)-1

	// Двигаемся с двух концов к центру
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// AreAnagrams проверяет, являются ли две строки анаграммами
// Временная сложность: O(n + m), где n и m - длины строк
// Пространственная сложность: O(k), где k - количество уникальных символов
func AreAnagrams(s1, s2 string) bool {
	// Если длины разные, то не анаграммы
	if len(s1) != len(s2) {
		return false
	}

	// Подсчитываем частоту символов
	charCount := make(map[rune]int)

	// Увеличиваем счетчик для первой строки
	for _, ch := range s1 {
		charCount[ch]++
	}

	// Уменьшаем счетчик для второй строки
	for _, ch := range s2 {
		charCount[ch]--
		if charCount[ch] < 0 {
			return false
		}
	}

	return true
}

// ReverseString переворачивает строку с учетом Unicode
// Временная сложность: O(n)
// Пространственная сложность: O(n)
func ReverseString(s string) string {
	// Преобразуем в слайс рун для корректной работы с Unicode
	runes := []rune(s)

	// Переворачиваем слайс
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// FirstUniqueChar находит первый уникальный символ в строке
// Временная сложность: O(n)
// Пространственная сложность: O(k), где k - количество уникальных символов
func FirstUniqueChar(s string) int {
	// Подсчитываем частоту каждого символа
	charCount := make(map[rune]int)

	for _, ch := range s {
		charCount[ch]++
	}

	// Находим первый символ с частотой 1
	for i, ch := range s {
		if charCount[ch] == 1 {
			return i
		}
	}

	return -1 // Не найден уникальный символ
}

// CharFrequency подсчитывает частоту символов в строке
// Временная сложность: O(n)
// Пространственная сложность: O(k), где k - количество уникальных символов
func CharFrequency(s string) map[rune]int {
	freq := make(map[rune]int)

	for _, ch := range s {
		freq[ch]++
	}

	return freq
}

// ============================================================================
// МАССИВЫ И СЛАЙСЫ
// ============================================================================

// TwoSum находит индексы двух чисел, сумма которых равна target
// Временная сложность: O(n)
// Пространственная сложность: O(n)
func TwoSum(nums []int, target int) []int {
	// Храним значение -> индекс
	numMap := make(map[int]int)

	for i, num := range nums {
		// Ищем дополнение
		complement := target - num

		if idx, found := numMap[complement]; found {
			return []int{idx, i}
		}

		// Сохраняем текущее число
		numMap[num] = i
	}

	return nil // Решение не найдено
}

// FindDuplicates находит все дубликаты в массиве
// Временная сложность: O(n)
// Пространственная сложность: O(n)
func FindDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	duplicates := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			duplicates[num] = true
		}
		seen[num] = true
	}

	// Преобразуем мапу в слайс
	result := make([]int, 0, len(duplicates))
	for num := range duplicates {
		result = append(result, num)
	}

	return result
}

// RotateArray сдвигает элементы массива вправо на k позиций
// Временная сложность: O(n)
// Пространственная сложность: O(1)
func RotateArray(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	// Нормализуем k
	k = k % n

	// Три переворота: [1,2,3,4,5,6,7], k=3
	// 1. [7,6,5,4,3,2,1]
	// 2. [5,6,7,4,3,2,1]
	// 3. [5,6,7,1,2,3,4]

	reverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	reverse(0, n-1)
	reverse(0, k-1)
	reverse(k, n-1)
}

// MergeSortedArrays объединяет два отсортированных массива
// Временная сложность: O(n + m)
// Пространственная сложность: O(n + m)
func MergeSortedArrays(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	// Сливаем массивы
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы
	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)

	return result
}

// MaxSubArraySum находит максимальную сумму подмассива (алгоритм Кадане)
// Временная сложность: O(n)
// Пространственная сложность: O(1)
func MaxSubArraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// Либо добавляем к текущей сумме, либо начинаем новую
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

// ============================================================================
// ЧИСЛА И МАТЕМАТИКА
// ============================================================================

// FizzBuzz классическая задача
// Временная сложность: O(n)
// Пространственная сложность: O(n)
func FizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i-1] = "FizzBuzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		case i%5 == 0:
			result[i-1] = "Buzz"
		default:
			result[i-1] = fmt.Sprintf("%d", i)
		}
	}

	return result
}

// IsPrime проверяет, является ли число простым
// Временная сложность: O(√n)
// Пространственная сложность: O(1)
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Проверяем делители вида 6k ± 1 до √n
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

// FactorialIterative вычисляет факториал итеративно
// Временная сложность: O(n)
// Пространственная сложность: O(1)
func FactorialIterative(n int) int {
	if n < 0 {
		return -1 // Ошибка
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

// FactorialRecursive вычисляет факториал рекурсивно
// Временная сложность: O(n)
// Пространственная сложность: O(n) - стек вызовов
func FactorialRecursive(n int) int {
	if n < 0 {
		return -1
	}
	if n <= 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

// FibonacciIterative вычисляет n-ое число Фибоначчи итеративно
// Временная сложность: O(n)
// Пространственная сложность: O(1)
func FibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	prev, curr := 0, 1

	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}

// FibonacciRecursive вычисляет n-ое число Фибоначчи рекурсивно
// Временная сложность: O(2^n) - неоптимально!
// Пространственная сложность: O(n)
func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// ReverseInteger переворачивает число
// Временная сложность: O(log n) - количество цифр
// Пространственная сложность: O(1)
func ReverseInteger(x int) int {
	result := 0

	for x != 0 {
		digit := x % 10
		x /= 10

		// Проверка на переполнение
		if result > math.MaxInt32/10 || result < math.MinInt32/10 {
			return 0
		}

		result = result*10 + digit
	}

	return result
}

// IsPowerOfTwo проверяет, является ли число степенью двойки
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func IsPowerOfTwo(n int) bool {
	// Степень двойки имеет только один бит равный 1
	// n & (n-1) обнуляет младший бит
	return n > 0 && (n&(n-1)) == 0
}

// ============================================================================
// СТРУКТУРЫ ДАННЫХ
// ============================================================================

// Stack реализация стека
type Stack struct {
	items []int
}

// Push добавляет элемент в стек - O(1)
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop извлекает элемент из стека - O(1)
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek возвращает верхний элемент без удаления - O(1)
func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty проверяет, пуст ли стек - O(1)
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Queue реализация очереди
type Queue struct {
	items []int
}

// Enqueue добавляет элемент в очередь - O(1)
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue извлекает элемент из очереди - O(n) из-за сдвига элементов
// Для O(1) можно использовать кольцевой буфер или два стека
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// IsEmpty проверяет, пуста ли очередь - O(1)
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// ListNode узел связного списка
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseLinkedList переворачивает связный список
// Временная сложность: O(n)
// Пространственная сложность: O(1)
func ReverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// HasCycle проверяет наличие цикла в связном списке (алгоритм Floyd)
// Временная сложность: O(n)
// Пространственная сложность: O(1)
func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// MiddleOfLinkedList находит середину связного списка
// Временная сложность: O(n)
// Пространственная сложность: O(1)
func MiddleOfLinkedList(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// IsValidParentheses проверяет валидность скобок
// Временная сложность: O(n)
// Пространственная сложность: O(n)
func IsValidParentheses(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// ============================================================================
// АЛГОРИТМЫ ПОИСКА И СОРТИРОВКИ
// ============================================================================

// BinarySearch бинарный поиск в отсортированном массиве
// Временная сложность: O(log n)
// Пространственная сложность: O(1)
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // Не найдено
}

// QuickSort быстрая сортировка
// Временная сложность: O(n log n) в среднем, O(n²) в худшем
// Пространственная сложность: O(log n) - стек рекурсии
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Копируем массив, чтобы не модифицировать оригинал
	result := make([]int, len(arr))
	copy(result, arr)

	quickSortHelper(result, 0, len(result)-1)
	return result
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pivotIdx := partition(arr, low, high)
		quickSortHelper(arr, low, pivotIdx-1)
		quickSortHelper(arr, pivotIdx+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// MergeSort сортировка слиянием
// Временная сложность: O(n log n)
// Пространственная сложность: O(n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// FindKthLargest находит k-й максимальный элемент (QuickSelect)
// Временная сложность: O(n) в среднем, O(n²) в худшем
// Пространственная сложность: O(1)
func FindKthLargest(nums []int, k int) int {
	// Ищем (len-k)-й элемент при сортировке по возрастанию
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	pivotIdx := partition(nums, left, right)

	if k == pivotIdx {
		return nums[k]
	} else if k < pivotIdx {
		return quickSelect(nums, left, pivotIdx-1, k)
	} else {
		return quickSelect(nums, pivotIdx+1, right, k)
	}
}

// ============================================================================
// РАБОТА С МАПАМИ
// ============================================================================

// CountFrequency подсчитывает частоту элементов
// Временная сложность: O(n)
// Пространственная сложность: O(k), где k - количество уникальных элементов
func CountFrequency(arr []int) map[int]int {
	freq := make(map[int]int)

	for _, num := range arr {
		freq[num]++
	}

	return freq
}

// GroupAnagrams группирует анаграммы
// Временная сложность: O(n * k log k), где n - количество строк, k - средняя длина строки
// Пространственная сложность: O(n * k)
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		// Сортируем символы строки как ключ
		key := sortString(s)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	runes := []rune(s)
	// Простая сортировка пузырьком для демонстрации
	// В продакшене лучше использовать sort.Slice
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

// ============================================================================
// РЕКУРСИЯ И ДЕРЕВЬЯ
// ============================================================================

// TreeNode узел бинарного дерева
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InOrderTraversal обход дерева In-Order (левый-корень-правый)
// Временная сложность: O(n)
// Пространственная сложность: O(h), где h - высота дерева
func InOrderTraversal(root *TreeNode) []int {
	result := []int{}

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		result = append(result, node.Val)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

// PreOrderTraversal обход дерева Pre-Order (корень-левый-правый)
// Временная сложность: O(n)
// Пространственная сложность: O(h)
func PreOrderTraversal(root *TreeNode) []int {
	result := []int{}

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

// PostOrderTraversal обход дерева Post-Order (левый-правый-корень)
// Временная сложность: O(n)
// Пространственная сложность: O(h)
func PostOrderTraversal(root *TreeNode) []int {
	result := []int{}

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		result = append(result, node.Val)
	}

	traverse(root)
	return result
}

// MaxDepth находит максимальную глубину бинарного дерева
// Временная сложность: O(n)
// Пространственная сложность: O(h)
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	return max(leftDepth, rightDepth) + 1
}

// Subsets генерирует все подмножества множества
// Временная сложность: O(2^n)
// Пространственная сложность: O(n * 2^n)
func Subsets(nums []int) [][]int {
	result := [][]int{}

	var backtrack func(start int, current []int)
	backtrack = func(start int, current []int) {
		// Добавляем копию текущего подмножества
		subset := make([]int, len(current))
		copy(subset, current)
		result = append(result, subset)

		// Генерируем все возможные продолжения
		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i+1, current)
			current = current[:len(current)-1]
		}
	}

	backtrack(0, []int{})
	return result
}

// ============================================================================
// ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ
// ============================================================================

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ============================================================================
// MAIN - ПРИМЕРЫ ИСПОЛЬЗОВАНИЯ
// ============================================================================

func main() {
	// Демонстрация некоторых функций

	fmt.Println("=== Строки ===")
	fmt.Println("IsPalindrome('racecar'):", IsPalindrome("racecar"))
	fmt.Println("AreAnagrams('listen', 'silent'):", AreAnagrams("listen", "silent"))
	fmt.Println("ReverseString('hello'):", ReverseString("hello"))

	fmt.Println("\n=== Массивы ===")
	fmt.Println("TwoSum([2,7,11,15], 9):", TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("MaxSubArraySum([-2,1,-3,4,-1,2,1,-5,4]):", MaxSubArraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	fmt.Println("\n=== Числа ===")
	fmt.Println("FizzBuzz(15):", FizzBuzz(15))
	fmt.Println("IsPrime(17):", IsPrime(17))
	fmt.Println("Fibonacci(10):", FibonacciIterative(10))

	fmt.Println("\n=== Сортировка ===")
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("QuickSort:", QuickSort(arr))
	fmt.Println("MergeSort:", MergeSort(arr))

	fmt.Println("\n=== Структуры данных ===")
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	val, _ := stack.Pop()
	fmt.Println("Stack Pop:", val)

	fmt.Println("IsValidParentheses('()[]{}'):", IsValidParentheses("()[]{}"))
}
