/*
Бинарный поиск
Реализовать алгоритм бинарного поиска встроенными методами языка.
Функция должна принимать отсортированный слайс и искомый элемент,
возвращать индекс элемента или -1, если элемент не найден.
Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.
*/
package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{8, 5, 2, 9, 10}

	fmt.Println(binarySearch(arr, 9))
}
