package main

import "fmt"

//4. WAP to Display Odd Numbers From 1 to 100
func main() {
	oddNumbers(100)
}

func oddNumbers(target int) {
	// wg := sync.WaitGroup{}
	for i := 0; i <= target; i++ {
		// wg.Add(1)
		// go func(num int) {
		// 	defer wg.Done()
		if i%2 != 0 {
			fmt.Printf("%d is odd\n", i)
		}
		// }(i)
	}
	// wg.Wait()

}
