/*
Удаление элемента слайса
Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.
Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:]))
и уменьшить длину слайса на 1.
*/
package main

func deleteElement[T any](slice []T, i int) []T {
	copy(slice[i:], slice[i+1:])

	var zero T
	slice[len(slice)-1] = zero

	return slice[:len(slice)-1]
}
