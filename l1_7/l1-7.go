/*
Конкурентная запись в map
Реализовать безопасную для конкуренции запись данных в структуру map.
Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
Проверьте работу кода на гонки (util go run -race).
*/
package main

import (
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[int]int),
	}
}

func (s *SafeMap) Set(key, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[key] = value
}

func (s *SafeMap) Get(key int) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.m[key]
	return val, ok
}

func main() {
	sm := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(i, i*i)
		}(i)
	}

	wg.Wait()
}
