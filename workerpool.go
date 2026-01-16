package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	workerPool := 5
	jobs := 10
	reqChan := make(chan int, jobs)
	resChan := make(chan int, jobs)
	for i := 0; i < workerPool; i++ {
		wg.Add(1)
		go worker(i, reqChan, resChan, wg)
	}
	for j := 0; j < jobs; j++ {

		reqChan <- j
	}
	close(reqChan)
	wg.Wait()
	close(resChan)
	for rData := range resChan {
		fmt.Println("response is :  ", rData)
	}
}

func worker(i int, reqChan chan int, resChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("i val ", i)
	for d := range reqChan {
		mu.Lock()
		resChan <- d * 2
		mu.Unlock()
	}
}
