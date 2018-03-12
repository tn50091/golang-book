package main

import (
	"fmt"
	"time"
	"sync"
)

type value struct {
	value int
	mu sync.Mutex
}

func main() {
	var a, b value
	var wg sync.WaitGroup
	wg.Add(2)
	go printSum(&a, &b, &wg)
	go printSum(&b, &a, &wg)
	wg.Wait() 
} 

func printSum(a, b *value, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mu.Lock()
	defer a.mu.Unlock()

	time.Sleep(2*time.Second)
	b.mu.Lock()
	defer b.mu.Unlock()

	fmt.Printf("sum=%v\n", a.value+b.value)
}