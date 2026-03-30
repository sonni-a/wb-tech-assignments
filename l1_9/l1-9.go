/*
Конвейер чисел
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
После этого данные из второго канала должны выводиться в stdout.
То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
Убедитесь, что чтение из второго канала корректно завершается.
*/
package main

import "fmt"

func generateNumbers(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func proccessNumbers(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()

	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	ch1 := generateNumbers(nums)
	ch2 := proccessNumbers(ch1)

	for result := range ch2 {
		fmt.Println(result)
	}
}
