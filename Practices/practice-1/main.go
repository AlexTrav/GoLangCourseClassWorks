package main

import "fmt"

// Практика - Урок - 6

// Sum - Суммирует и возвращает сумму всех элементов исходного слайса
func Sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// FilterEven - Возвращает слайс чётных чисел из исходного слайса
func FilterEven(nums []int) []int {
	result := make([]int, 0, len(nums))
	for _, v := range nums {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

// Reverse - Разворачивает слайс
func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// Reverse2 - Разворачивает слайс
func Reverse2(nums []int) []int {
	result := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		result = append(result, nums[i])
	}
	return result
}

// Remove — удаляет элемент из слайса по указанному индексу и возвращает новый слайс без этого элемента
func Remove(nums []int, index int) []int {
	if index < 0 || index >= len(nums) {
		return nums
	}
	return append(nums[:index], nums[index+1:]...)
}

// Remove2 — удаляет элемент из слайса по указанному индексу. Возвращает новый слайс
func Remove2(nums []int, index int) []int {
	if index < 0 || index >= len(nums) {
		return nums
	}

	result := make([]int, 0, len(nums)-1)

	for i := 0; i < len(nums); i++ {
		if i == index {
			continue
		}
		result = append(result, nums[i])
	}

	return result
}

// Remove3 — удаляет элемент из слайса по индексу. Возвращает новый слайс
func Remove3(nums []int, index int) []int {
	if index < 0 || index >= len(nums) {
		return nums
	}
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	return nums[:len(nums)-1]
}

// Insert - Вставляет элемент по индексу
func Insert(nums []int, index int, value int) []int {
	if index < 0 {
		index = 0
	}
	if index > len(nums) {
		index = len(nums)
	}

	nums = append(nums, 0)
	copy(nums[index+1:], nums[index:])
	nums[index] = value

	return nums
}

// Chunk - Разбивает слайс nums на чанки фиксированного размера chunkSize.
func Chunk(nums []int, chunkSize int) [][]int {
	if chunkSize <= 0 {
		return nil
	}

	chunksCount := (len(nums) + chunkSize - 1) / chunkSize
	result := make([][]int, 0, chunksCount)

	for i := 0; i < len(nums); i += chunkSize {
		end := i + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		result = append(result, nums[i:end])
	}
	return result
}

func main() {
	s := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Printf("Исходный слайс: %d\n", s)

	fmt.Printf("Сумма элементов слайса: %d\n", Sum(s))

	fmt.Printf("Cлайс чётных чисел: %d\n", FilterEven(s))

	Reverse(s)
	fmt.Printf("Слайс в обратном порядке: %d\n", s)

	fmt.Printf("Слайс в обратном порядке 2: %d\n", Reverse2(s)) // Не меняет исходный слайс

	// Вызываем функции на копиях
	s1 := append([]int(nil), s...)
	fmt.Printf("Слайс без удалённого элемента под индексом 2 Remove: %d\n", Remove(s1, 2))

	s2 := append([]int(nil), s...)
	fmt.Printf("Слайс без удалённого элемента под индексом 2 Remove2: %d\n", Remove2(s2, 2)) // версия без подслайсов

	s3 := append([]int(nil), s...)
	fmt.Printf("Слайс без удалённого элемента под индексом 2 Remove3: %d\n", Remove3(s3, 2)) // версия со сдвигом

	s4 := append([]int(nil), s...)
	inserted := Insert(s4, 3, 999)
	fmt.Printf("Слайс после вставки 999 по индексу 3: %d\n", inserted)

	chunks := Chunk(s, 3)
	fmt.Printf("Слайс, разбитый на чанки по 3 элемента: %d\n", chunks)
}

/* Вывод из консоли:
Исходный слайс: [10 9 8 7 6 5 4 3 2 1]
Сумма элементов слайса: 55
Cлайс чётных чисел: [10 8 6 4 2]
Слайс в обратном порядке: [1 2 3 4 5 6 7 8 9 10]
Слайс в обратном порядке 2: [10 9 8 7 6 5 4 3 2 1]
Слайс без удалённого элемента под индексом 2 Remove: [1 2 4 5 6 7 8 9 10]
Слайс без удалённого элемента под индексом 2 Remove2: [1 2 4 5 6 7 8 9 10]
Слайс без удалённого элемента под индексом 2 Remove3: [1 2 4 5 6 7 8 9 10]
Слайс после вставки 999 по индексу 3: [1 2 3 999 4 5 6 7 8 9 10]
Слайс, разбитый на чанки по 3 элемента: [[1 2 3] [4 5 6] [7 8 9] [10]]
*/
