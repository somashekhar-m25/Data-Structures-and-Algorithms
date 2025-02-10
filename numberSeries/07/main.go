package main

import "fmt"

//7. WAP to Generate Multiplication Table
func main() {
	MultipliactionTables(2)
}

func MultipliactionTables(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d X %d = %d\n", num, i, num*i)
	}
}
