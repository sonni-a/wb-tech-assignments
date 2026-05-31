/*
Обмен значениями без третьей переменной
Поменять местами два числа без использования временной переменной.

Подсказка: примените сложение/вычитание или XOR-обмен.
*/
package main

import "fmt"

func swapSum(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func swapXOR(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	x, y := swapSum(a, b)
	fmt.Println(x, y)

	x, y = swapXOR(a, b)
	fmt.Println(x, y)
}
