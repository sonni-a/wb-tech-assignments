/*
Конкурентное возведение в квадрат
Написать программу, которая конкурентно рассчитает значения квадратов чисел,
взятых из массива [2,4,6,8,10], и выведет результаты в stdout.
Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	res := make(chan int, len(nums))

	for _, n := range nums {
		wg.Add(1)

		go func(num int) {
			defer wg.Done()
			res <- num * num
		}(n)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		fmt.Println(r)
	}
}
