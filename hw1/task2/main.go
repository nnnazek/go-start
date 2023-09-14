package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var radius float64

	fmt.Print("Print radius: ")

	fmt.Fscan(os.Stdin, &radius)

	area := math.Pi * math.Pow(radius, 2)

	circumference := 2 * math.Pi * radius

	fmt.Printf("Area of circle: %.2f\n", area)
	fmt.Printf("Circumference: %.2f\n", circumference)
}
