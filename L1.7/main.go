package main

import (
	"fmt"
	"sync"
)

const (
	goroutines = 100
	increments = 1000
)

type SafeMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{m: make(map[K]V)}
}

func (s *SafeMap[K, V]) Get(k K) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m[k]
	return v, ok
}

func (s *SafeMap[K, V]) Set(k K, v V) {
	s.mu.Lock()
	s.m[k] = v
	s.mu.Unlock()
}

func (s *SafeMap[K, V]) Update(k K, f func(old V) V) {
	s.mu.Lock()
	s.m[k] = f(s.m[k])
	s.mu.Unlock()
}

func main() {
	sm := NewSafeMap[string, int]()
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				sm.Update("counter", func(old int) int { return old + 1 })
			}
		}()
	}

	wg.Wait()
	v, _ := sm.Get("counter")
	fmt.Println("counter =", v) // ожидание: 100 * 1000 = 100000
}
