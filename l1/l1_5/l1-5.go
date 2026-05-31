/*
Таймаут на канал
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала – читать эти значения.
По истечении N секунд программа должна завершаться.
Подсказка: используйте time.After или таймер для ограничения времени работы.
*/
package main

import (
	"fmt"
	"time"
)

func timeoutChannel(n time.Duration) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		i := 0
		timeout := time.After(n)

		for {
			select {
			case ch <- i:
				i++
			case <-timeout:
				return
			}
		}
	}()

	return ch
}

func main() {
	var n int
	fmt.Println("Введите количество секунд:")
	fmt.Scan(&n)

	ch := timeoutChannel(time.Duration(n) * time.Second)

	for v := range ch {
		fmt.Println(v)
	}
}
