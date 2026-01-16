// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 2)
	//	ch1 := make(chan int)
	wg.Add(2)
	go add(2, 4, ch)
	go add(2, 6, ch)
	//for i := 0; i < 2; i++ {
	//select {
	d := <-ch
	fmt.Println("first ", d)
	//case d1 := <-ch:
	//	fmt.Println("first ", d1)
	//default:
	//	fmt.Println("here")

	//	}
	//}
	wg.Wait()
}

func add(a, b int, c chan int) {
	defer wg.Done()
	d := a + b
	c <- d

}
