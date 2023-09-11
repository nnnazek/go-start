package main

import (
	"fmt"
	"os"
	"task1/operations"
)

func main() {
	var number1 int
	var number2 int
	fmt.Print("Print the first number: ")
	fmt.Fscan(os.Stdin, &number1)

	fmt.Print("Print the second number: ")
	fmt.Fscan(os.Stdin, &number2)

	fmt.Println(operations.Add(number1, number2))
}
