/*
Быстрая сортировка (quicksort)
Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.
Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
Для выбора опорного элемента можно взять середину или первый элемент.
*/
package main

import "fmt"

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivotIndex := partition(nums, left, right)

	quickSort(nums, left, pivotIndex-1)
	quickSort(nums, pivotIndex+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left - 1

	for j := left; j < right; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

func main() {
	arr := []int{8, 5, 2, 9, 10}
	QuickSort(arr)
	fmt.Println(arr)
}
