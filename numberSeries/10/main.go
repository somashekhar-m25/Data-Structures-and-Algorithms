package main

import "fmt"

//10. How to Reverse a Number?
func main() {
	ReverseNumber(12345)
}

func ReverseNumber(num int) {
	fmt.Printf("num before revrse: %d\n", num)
	reverseNum := 0
	for {
		if num < 10 {
			reverseNum = reverseNum*10 + num
			break
		}
		reverseNum = reverseNum*10 + num%10
		num = num / 10
	}
	fmt.Printf("num after revrse: %d\n", reverseNum)
}
