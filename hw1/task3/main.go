package main

import (
	"fmt"
	"os"
)

func Check(num float64) bool {
	return int(num)%2 == 0
}

func printResult(num float64) {
	if Check(num) {
		fmt.Println("THE NUMBER", num, "IS AN EVEN NUMBER")
	} else {
		fmt.Println("THE NUMBER", num, "IS AN ODD NUMBER")
	}
}

func main() {
	var number float64

	fmt.Print("Print number: ")
	fmt.Fscan(os.Stdin, &number)

	printResult(number)
}
