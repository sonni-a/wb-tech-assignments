/*
Пересечение множеств
Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) —
т.е. вывести элементы, присутствующие и в первом, и во втором.
Пример:
A = {1,2,3}
B = {2,3,4}
Пересечение = {2,3}
*/
package main

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	result := []int{}

	for _, num := range nums1 {
		set[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			result = append(result, num)
			delete(set, num)
		}
	}

	return result
}
