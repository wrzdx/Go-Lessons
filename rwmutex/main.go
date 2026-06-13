package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Storage interface {
	Set(key, value string)
	Get(key string) (string, bool)
}

type MutexStorage struct {
	mu sync.Mutex
	m  map[string]string
}

func NewMutexStorage() *MutexStorage {
	return &MutexStorage{
		m: make(map[string]string),
	}
}

func (s *MutexStorage) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[key] = value
}

func (s *MutexStorage) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.m[key]
	return value, ok
}

type RWMutexStorage struct {
	mu sync.RWMutex
	m  map[string]string
}

func NewRWMutexStorage() *RWMutexStorage {
	return &RWMutexStorage{
		m: make(map[string]string),
	}
}

func (s *RWMutexStorage) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[key] = value
}

func (s *RWMutexStorage) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.m[key]
	return value, ok
}

func benchmark(storage Storage) time.Duration {
	start := time.Now()

	var wg sync.WaitGroup

	// писатели
	for i := 0; i < 10; i++ {
		i := i

		wg.Go(func() {
			for j := 0; j < 100_000; j++ {
				key := strconv.Itoa(i) + "-" + strconv.Itoa(j)

				storage.Set(key, "value")
			}
		})
	}

	// читатели
	for i := 0; i < 100; i++ {
		wg.Go(func() {
			for j := 0; j < 100_000; j++ {
				key := strconv.Itoa(j % 1000)

				storage.Get(key)
			}
		})
	}

	wg.Wait()

	return time.Since(start)
}

func main() {
	mutexStorage := NewMutexStorage()
	rwStorage := NewRWMutexStorage()

	mutexTime := benchmark(mutexStorage)
	rwTime := benchmark(rwStorage)

	fmt.Println("Mutex:", mutexTime)
	fmt.Println("RWMutex:", rwTime)
}
