package main

import "fmt"

//6. WAP to Find Largest of Three Numbers
func main() {
	fmt.Println("Given array : ", []int{12, 45, 23})
	fmt.Printf("largest number in the given array is : %d", LargestOfThree([]int{12, 45, 23}))
}

func LargestOfThree(numbers []int) int {
	largest := 0
	for _, v := range numbers {
		if largest < v {
			largest = v
		}
	}
	return largest
}
