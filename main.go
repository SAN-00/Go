package main

import (
	"fmt"
	"myproject/simplecalc"
)

// --- Structures ---
type Movie struct {
	Title string
	Hero  string
	Year  int
}

func main() {
	favMovie := Movie{
		Title: "Baahubali",
		Hero:  "Prabhas",
		Year:  2015,
	}

	fmt.Println("Title:", favMovie.Title)
	fmt.Println("Hero:", favMovie.Hero)
	fmt.Println("Year:", favMovie.Year)

	// --- Array ---

	var numbers [5]int
	numbers[0] = 5
	numbers[1] = 10
	numbers[2] = 15
	numbers[3] = 20
	numbers[4] = 25

	fmt.Println("Array values:")
	for i, v := range numbers {
		fmt.Println(i, v)
	}

	fruits := [5]string{"Apple", "Banana", "Cherry", "Mango", "Kiwi"}
	fmt.Println("Fruits:", fruits)

	// ----------- Slices -----------

	weekdays := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	slice1 := weekdays[1:4]
	fmt.Println("Slice1:", slice1)

	slice2 := []int{100, 200, 300}
	fmt.Println("Slice2 before append:", slice2)
	slice2 = append(slice2, 400, 500)
	fmt.Println("Slice2 after append:", slice2)

	fmt.Println("Length of slice2:", len(slice2))
	fmt.Println("Capacity of slice2:", cap(slice2))

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

	// --- defer ---
	defer fmt.Println("Second")
	defer fmt.Println("Middle")
	defer fmt.Println("First")
}

func increase(num *int) {
	*num = *num + 1
}
