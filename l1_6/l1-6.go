/*
Остановка горутины
Реализовать все возможные способы остановки выполнения горутины.
Классические подходы: выход по условию, через канал уведомления, через контекст,
прекращение работы runtime.Goexit() и др.
Продемонстрируйте каждый способ в отдельном фрагменте кода.
*/
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Exiting goroutine by condition:")
	doneByCondition()

	fmt.Println("Exiting goroutine by channel:")
	doneByChannel()

	fmt.Println("Exiting goroutine by context:")
	doneByContext()

	fmt.Println("Exiting goroutine by Goexit():")
	doneByGoExit()
}

// выход по условию
func doneByCondition() {
	var wg sync.WaitGroup
	done := false
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if done {
				fmt.Println("goroutine stopped by condition!")
				return
			}
			fmt.Println("working...", i)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(3 * time.Second)
	done = true
	wg.Wait()
}

// выход через канал уведомления
func doneByChannel() {
	var wg sync.WaitGroup
	doneCh := make(chan struct{})
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-doneCh:
				fmt.Println("goroutine stopped by channel!")
				return
			default:
				fmt.Println("working...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	close(doneCh)
	wg.Wait()
}

// выход через контекст
func doneByContext() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine stopped by context!")
				return
			default:
				fmt.Println("working...")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}

// выход через прекращение работы runtime.Goexit()
func doneByGoExit() {
	go func() {
		fmt.Println("goroutine started!")
		runtime.Goexit()
	}()
	time.Sleep(time.Second)
}
