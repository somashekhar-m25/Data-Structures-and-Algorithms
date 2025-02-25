package main

import "fmt"

//15.Disarium number Sum of its digits powered with their respective position is equal to the original number.

func main() {
	DisariumNumber(89)
}

func DisariumNumber(num int) {
	temp := num
	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	num = temp
	sum := 0
	for num > 0 {
		sum += PowerOf(num%10, count)
		num /= 10
		count--
	}
	if sum == temp {
		fmt.Printf("is Disarium number: %d\n", temp)
	} else {
		fmt.Printf("is not Disarium number: %d", temp)
	}

}

func PowerOf(num, power int) int {
	v := 1
	for i := 1; i <= power; i++ {
		v *= num
	}
	return v
}
