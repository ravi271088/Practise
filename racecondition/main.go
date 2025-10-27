package main

import (
	"fmt"
	"sync"
)

var count int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	wg.Add(2)
	go counter()
	go counter()
	wg.Wait()
	fmt.Println(count)
}

func counter() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
}
