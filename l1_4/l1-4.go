/*
Завершение по Ctrl+C
Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.
Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.
*/
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d working...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, wg)
	}

	<-sigs
	fmt.Println("Ctrl+C pressed, stopping workers...")
	cancel()

	wg.Wait()
	fmt.Println("All workers stopped, exiting")
}
