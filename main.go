package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, Welcome to GoLang with Eziz")

	// Variables and Constants
	isbool := true
	hate := false
	win := true
	var Name string = "Eziz Soyunov"
	var a int = 5
	var b float32 = 4.32
	const pi float64 = 3.14159265359
	x, y := 5, 6

	// Menu
	fmt.Println("1. Hello World : ")
	fmt.Println("2. Arithmetic Operators: ")
	fmt.Println("3. Logical Operators:  ")
	fmt.Println("4. Printf:  ")
	fmt.Println("5. Pointers: ")
	fmt.Println("6. Loops: ")
	fmt.Println("7. Desicion Making: ")
	fmt.Println("8. Arrays: ")
	fmt.Println("9. Maps: ")
	fmt.Println("10. Functions : ")
	fmt.Println("11. Recursion : ")
	fmt.Println("12. Defer,Recover, Panic : ")
	fmt.Println("13. Structures and Interfaces : ")
	fmt.Println("14. File Input and Output : ")
	fmt.Println("15. Webserver : ")
	fmt.Println("16. GoLang, MySql tutorial: ")
	fmt.Print("Please choose one of them: ")

	// var then variable name then variable type
	var first string
	fmt.Scanln(&first)
	x, error := strconv.Atoi(first)
	fmt.Println(error)

	switch x {
	case 1:
		fmt.Println(" 1. Hello World ")
		// get string length

	case 2:
		fmt.Println("2. Arithmetic Operators:")
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(pi)
		fmt.Println("Arithmetic operation")
		fmt.Println("Arithmetic addtion")
		fmt.Println("x+y= ", x+y)
		fmt.Println("Arithmetic subtraction")
		fmt.Println("x-y= ", x-y)
		fmt.Println("Arithmetic multiplication")
		fmt.Println("x*y= ", x*y)
		fmt.Println("Arithmetic division")
		fmt.Println("x/y= ", x/y)
		fmt.Println("Arithmetic modulus")
		fmt.Println("x mod y= ", x%y)
	case 3:
		fmt.Println("3. Logical Operators: ")
		fmt.Println("&& -and example", isbool && hate)
		fmt.Println("|| -or example", isbool || hate)
		fmt.Println("! negation example ", !isbool, !hate)
	case 4:
		fmt.Println("4. Printf:")

		fmt.Printf(" %T \n", isbool)

		fmt.Println("String length=", len(Name))
		fmt.Println(Name + "is a GoLang developer")

		fmt.Printf(" %.3f \n", pi)
		fmt.Printf(" %T \n", Name)
		fmt.Printf(" %t \n", win)
		fmt.Printf(" %d \n", x)
		fmt.Printf(" %b \n", 25)
		fmt.Printf(" %c \n", 34)
		fmt.Printf(" %x \n", 15)
	case 5:
		//Pointer example function
		fmt.Println("5. Pointers:")
		changeValue(x)
		fmt.Println("Without pointer", x)
		changeValueP(&x)
		fmt.Println("With Pointer", x)
	case 6:
		fmt.Println("6.Loops:")
		for i := 1; i < x; i++ {
			for j := 1; j < i; j++ {
				fmt.Printf("*_*")
			}
			fmt.Println(i)
		}
	case 7:
		fmt.Println("7. Decision Making:")
		fmt.Print("please enter your age: ")
		var age string
		fmt.Scanln(&age)
		yas, error := strconv.Atoi(age)
		fmt.Println(error)

		if yas > 18 {
			fmt.Println("Yes, you can vote")
		} else {
			fmt.Println("No, you can't vote")
		}

	case 8:
		fmt.Println("8. Arrays ")
		fmt.Print("please enter quantity of students: ")
		var quantity string
		var sany int
		var Students [100]int
		fmt.Scanln(&quantity)
		sany, error := strconv.Atoi(quantity)
		fmt.Println(error)

		for n := 0; n < sany; n++ {
			Students[n] = n + 1
			fmt.Println(Students[n])

		}

		EvenNum := [5]int{0, 2, 4, 6, 8}
		for i, value := range EvenNum {
			fmt.Println(value, i)
		}

		numSlice := []int{5, 4, 3, 2, 1}
		sliced := numSlice[0:]
		fmt.Println(sliced)
		slice2 := make([]int, 5, 10)
		copy(slice2, numSlice)
		fmt.Println(numSlice)

		slice3 := append(numSlice, 3, 0, -1)
		fmt.Println(slice3)

	case 9:
		{
			// we choose map example
			fmt.Println("9. Maps - Welcome to Maps examples")
			// first map example
			Talyp := make(map[string]int)
			Talyp["Mahym"] = 15
			Talyp["Myrat"] = 13
			Talyp["Sahet"] = 12
			Talyp["Kerwen"] = 17

			fmt.Println(Talyp["Mahym"])
			fmt.Println(len(Talyp))

			Talyp["Mahriban"] = 14
			fmt.Println(len(Talyp))
			delete(Talyp, "Mahriban")
			fmt.Println(len(Talyp))
			// second  map example

			superhero :=
				map[string]map[string]string{
					"Superman": map[string]string{
						"Realname": "Clark Kent",
						"City":     "Mertopolis",
					},

					"Batman": map[string]string{
						"Realname": "Bruce Wayne",
						"City":     "Gotham City",
					},
				}
			if temp, hero := superhero["Batman"]; hero {
				fmt.Println(temp["Realname"], temp["City"])
			}

		}
	case 10:
		// simple function example
		fmt.Println("10.Welcome to Functions examples ")

		x, y := 5, 6
		fmt.Println(add(x, y))
	case 11:
		// simple Recursion factorial example
		fmt.Println("11. Welcome to Recursion example ")
		num := 5
		fmt.Println(num, "number factorial is ", factorial(num))
	case 12:
		fmt.Println("12. Defer, Recover, Panic ")
		// defer example Lifo
		defer FirstRun()
		SecondRun()

		// Recover example
		fmt.Println(div(3, 0))
		fmt.Println(div(3, 5))
		// Panic example
		demPanic()

		fmt.Println(addemup(10, 20, 30, 40, 50))
	case 13:
		fmt.Println("13. Welcome to Structures and Interfaces ")
		// Structure example
		rect := Rectangle{50, 60}
		circ := Circle{7}

		fmt.Println(rect.height, rect.width)

		// we need to calculate rectangle area
		fmt.Println("Area of rectangle is ", rect.area())
		// we need to calculate diognal of rectangle
		fmt.Println("Diognal of rectangle is ", rect.diognal())

		// Interface example
		fmt.Println("Area of rectangle is ", getArea(rect))
		fmt.Println("Area of Circle ", getArea(circ))

	case 14:
		fmt.Println("14. Welcome to File input and Output ")
		file, err := os.Create("sample.txt")

		if err != nil {
			log.Fatal(err)
		}

		file.WriteString("Hi, My name iz Eziz and this file crated using GoLang")

		file.Close()

		stream, err1 := ioutil.ReadFile("sample.txt")

		if err1 != nil {
			log.Fatal(err1)
		}

		s1 := string(stream)
		fmt.Println(s1)

	case 15:
		fmt.Println("15. Welcome to Web Server ")
		// lets do web server example by using Golang

		http.HandleFunc("/", handler)
		http.HandleFunc("/Hello", handler2)
		http.ListenAndServe(":8080", nil)

	case 16:

		fmt.Println("16. Welcome to MySql totorial ")
		// lets do connection with MySql example by using Golang

		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		fmt.Println(" Successfully Connected to MySql database")

	default:
		fmt.Println("you didnt choose anything")
	}
}

// simple add function

func add(x, y int) int {
	return x + y
}

func changeValue(x int) {
	x = 7
}

func changeValueP(x *int) {
	*x = 7
}

// simple factorial function

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

// simple defer, panic function
func FirstRun() {
	fmt.Println("I executed first ")
}

func SecondRun() {
	fmt.Println("I executed second ")
}

// simple recover, panic function
func div(num1, num2 int) int {

	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

// simple panic function
func demPanic() {
	fmt.Println(recover())
}

func addemup(args ...int) int {

	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}

// simple structure example

type Rectangle struct {
	height float64
	width  float64
}

// We calculate rectangle area by using rectangle structure
func (rect Rectangle) area() float64 {

	return rect.height * rect.width
}

// We calculate rectangle diognal by using rectangle structure
func (rect Rectangle) diognal() float64 {

	dio := math.Sqrt(rect.height*rect.height + rect.width*rect.width)
	return dio
}

// how to use interface, i will show you by example

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c1 Circle) area() float64 {

	return math.Pi * math.Pow(c1.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "welcome to inceakym.com")
}
func handler2(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "welcome to bildirisyeri.com")
}
