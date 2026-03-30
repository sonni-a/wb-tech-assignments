/*
Установка бита в числе
Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
Подсказка: используйте битовые операции (|, &^).
*/
package main

import "fmt"

func SetBit(num int64, i int, val int) (int64, error) {
	if val != 0 && val != 1 {
		return num, fmt.Errorf("bit value must be 0 or 1")
	}

	if val == 0 {
		num &^= (1 << i)
	} else {
		num |= (1 << i)
	}

	return num, nil
}

func main() {
	var num int64
	fmt.Println("Введите число: ")
	fmt.Scan(&num)

	var i int
	fmt.Println("Введите номер бита: ")
	fmt.Scan(&i)

	var val int
	fmt.Println("Введите значение бита (0 или 1): ")
	fmt.Scan(&val)

	result, err := SetBit(num, i, val)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
