package main

import (
	"fmt"
	"myproject/simplecalc"
	"strings"
	"sync"
)

// --- Palindrome Function ---
func Palindrome(word string) {
	length := len(word)
	for i := 0; i < length/2; i++ {
		if word[i] != word[length-1-i] {
			fmt.Println(word, "is NOT a palindrome")
			return
		}
	}
	fmt.Println(word, "is a palindrome")
}

// --- Fibonacci Sequence ---
func Fibonacci(limit int) {
	a, b := 0, 1
	fmt.Println("Fibonacci Sequence:")
	for a <= limit {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

// --- HelloGreetings in Different Languages ---
func greet(lang, msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s: %s\n", lang, msg)
}

func HelloGreetings() {
	fmt.Println("\nGreetings in different languages:")
	var wg sync.WaitGroup
	wg.Add(5)
	go greet("English", "Hello!", &wg)
	go greet("Spanish", "Hola!", &wg)
	go greet("French", "Bonjour!", &wg)
	go greet("Tamil", "Vanakkam!", &wg)
	go greet("Telugu", "Namaskaram!", &wg)

	fmt.Scanln()
}

// --- Factorial Program ---

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// --- Counting Words in a Sentence ---
func wordCounter(sentence string) {
	words := strings.Split(sentence, " ")

	fmt.Println("\nWords and their lengths:")
	for _, word := range words {
		fmt.Printf("%s %d\n", word, len(word))
	}
}

type Student struct {
	Name  string
	Marks int
}

func (s *Student) ShowResult() string {
	if s.Marks >= 40 {
		return s.Name + " has Passed"
	}
	return s.Name + " has Failed"
}

// --- Pointer function ---
func IncreaseMarks(marks *int) {
	*marks = *marks + 10
}

func main() {
	fmt.Println("Palindrome Check:")
	Palindrome("level")
	Palindrome("sai")

	fmt.Println("\nFibonacci till 30:")
	Fibonacci(30)

	HelloGreetings()

	fmt.Println("\nFactorial Calculation:")
	fmt.Println("Factorial of 5:", Factorial(5))
	fmt.Println("Factorial of 7:", Factorial(7))

	wordCounter("My name is Sai Abhishiktha")

	// --- Maps ---
	subjects := map[string]int{
		"Maths":   75,
		"Science": 60,
		"English": 50,
	}

	fmt.Println("\nTotal Subjects:", len(subjects))

	fmt.Println("Subjects:", subjects)

	// Adding new subject
	subjects["History"] = 80
	fmt.Println("After adding History:", subjects)

	// Updating marks
	subjects["Maths"] = 85
	fmt.Println("After updating Maths:", subjects)

	// Checking for a subject
	if val, ok := subjects["Science"]; ok {
		fmt.Println("Science marks found:", val)
	} else {
		fmt.Println("Science not found")
	}

	fmt.Println("All Subjects and Marks:")
	for sub, mark := range subjects {
		fmt.Println(sub, ":", mark)
	}

	// --- Pointer function ---
	score := 35
	fmt.Println("\nBefore increase:", score)
	IncreaseMarks(&score)
	fmt.Println("After increase:", score)

	// --- Method ---
	stu1 := Student{Name: "Sammy", Marks: 55}
	stu2 := Student{Name: "Abby", Marks: 30}

	fmt.Println(stu1.Name, "Result:", stu1.ShowResult())
	fmt.Println(stu2.Name, "Result:", stu2.ShowResult())
	mainMovie()
}

// --- Assgn 2 ---
// --- Structures ---
type Movie struct {
	Title string
	Hero  string
	Year  int
}

func mainMovie() {
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

	// --- Assgn1 ---

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
