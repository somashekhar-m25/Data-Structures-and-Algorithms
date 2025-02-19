package main

import (
	"fmt"
)

// 3. WAP to Display Even Numbers From 1 to 100s
func main() {
	EvenNumbers(100)
}
func EvenNumbers(target int) {
	// wg := sync.WaitGroup{}
	for i := 0; i <= target; i++ {
		// wg.Add(1)
		// go func(num int) {
		// 	defer wg.Done()
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		}
		// }(i)
	}
	// wg.Wait()

}
