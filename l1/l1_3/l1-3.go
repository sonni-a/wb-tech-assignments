/*
Работа нескольких воркеров
Реализовать постоянную запись данных в канал (в главной горутине).
Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Job struct {
	ID int
}

func Worker(id int, jobs <-chan Job) {
	for job := range jobs {
		fmt.Printf("Worker %d processing Job ID: %d\n", id, job.ID)
	}
}

func main() {
	numWorkers := flag.Int("n", 3, "number of workers")
	flag.Parse()

	jobs := make(chan Job)

	for i := 1; i <= *numWorkers; i++ {
		go Worker(i, jobs)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for i := 1; ; i++ {
		select {
		case jobs <- Job{ID: i}:
		case <-sigChan:
			fmt.Println("\nShutting down...")
			close(jobs)
			return
		}
	}
}
