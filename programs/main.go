package main

import "fmt"

/*func checkPalidrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}

	}
	return true
}

func main() {
	println(checkPalidrome("racecar")) // true
	println(checkPalidrome("hello"))   // false
	println(reverseString("hello"))    // "olleh"
	println(reverseString("racecar"))  // "racecar"

}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
*/
// Goroutine function to generate Fibonacci series
/*
func generateFibonacci(n int, ch chan []int) {
	fib := make([]int, n)
	if n > 0 {
		fib[0] = 0
	}
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	ch <- fib // Send result to channel
}

func main() {
	n := 10
	ch := make(chan []int)
	ch1 := make(chan []int)

	go generateFibonacci(n, ch)  // Launch goroutine
	go generateFibonacci(5, ch1) // Launch goroutine

	result := <-ch // Receive result from channel
	fmt.Println("Fibonacci series:", result)

	result1 := <-ch1 // Receive result from channel
	fmt.Println("Fibonacci seriesss:", result1)
}

func uniqueNames(a, b []string) []string {
	nameSet := make(map[string]int)
	for _, name := range a {
		nameSet[name] = 1
	}
	fmt.Printf("ff %+v\n", nameSet)
	for _, name := range b {
		nameSet[name] = 2
	}

	fmt.Printf("ff %+v\n", nameSet)

	result := make([]string, 0, len(nameSet))
	for i, _ := range nameSet {
		result = append(result, i)
	}

	return result
}

func main() {
	a := []string{"Ava", "Emma", "Olivia"}
	b := []string{"Olivia", "Sophia", "Emma"}
	fmt.Println(uniqueNames(a, b))
}
*/
func extractFirstChar(strFirstChar string) string {
	if len(strFirstChar) < 1 {
		return ""
	}
	return string([]rune(strFirstChar)[0:1])
}

func main() {
	str := extractFirstChar("xbcdef")
	fmt.Println(str)
}
