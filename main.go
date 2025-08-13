package main

import (
	"fmt"
	"myproject/simplecalc"
)

func main() {
	fmt.Println("Hello, Sai")

	a, b := 6, 4
	fmt.Println(simplecalc.Add(a, b))

	c, d := 10, 3
	fmt.Println(simplecalc.Sub(c, d))

	e, f := 20, 4
	fmt.Println(simplecalc.Div(e, f))

	x, y := 30, 3
	fmt.Println(simplecalc.Mul(x, y))

	// --- if / else if / else ---
	num := 24
	if num > 24 {
		fmt.Println("Number is greater than 24")
	} else if num == 24 {
		fmt.Println("Number is exactly 24")
	} else {
		fmt.Println("Number is less than 24")
	}

	// --- switch case ---
	day := "Wednesday"
	switch day {
	case "Wednesday":
		fmt.Println("Weekday")
	case "Saturday":
		fmt.Println("Weekend")
	default:
		fmt.Println("It's another day")
	}

	// --- for loop ---
	fmt.Println("Counting from 1 to 7:")
	for i := 1; i <= 7; i++ {
		fmt.Println(i)
	}
	// --- pointers ---
	j := 5
	fmt.Println("Before:", j)
	increase(&j)
	fmt.Println("After:", j)

	// --- defer  ---

	defer fmt.Println("Second")
	defer fmt.Println("Middle")
	defer fmt.Println("First")

}

func increase(num *int) {
	*num = *num + 1
}
