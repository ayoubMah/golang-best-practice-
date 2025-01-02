package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	// ----------------------- For ------------------------------------------------
	/*
		The basic for loop has three components separated by semicolons:

		the init statement: executed before the first iteration
		the condition expression: evaluated before every iteration
		the post statement: executed at the end of every iteration
	*/

	//var sum int = 0

	for i := 0; i <= 5; i++ {
		fmt.Printf("The Element number: %v \n", i)
	}

	fmt.Println("____________________")

	// the init and the post statements are optional
	i := 0
	for i <= 5 {
		i++
		fmt.Printf("The Element number: %v \n", i)
	}

	// now you can delete the semi colones and you gana while loop
	j := 1
	for j < 1000 {
		j += j
	}
	fmt.Println(j)
	fmt.Println("___________________")

	fmt.Println(sqrt(16))
	fmt.Println(sqrt(-16))

	// ----------------------- If ------------------------------------------------

	fmt.Println(
		pow(2, 3, 2),
		pow(33, 4, 10000),
	)

	// ----------------------- Switch --------------------------------------------

	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Print(" Linux, Omakub lol \n")
	case "darwin":
		fmt.Println("OS X")
	}

	fmt.Println("Is Friday Far ?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today")

	case today + 1:
		fmt.Println("Tomorrow")

	case today + 2:
		fmt.Println("In two days")

	default:
		fmt.Println("too far away")
	}

	// A defer statement defers the execution of a function until the surrounding function returns.
	defer fmt.Println("Bye")

	fmt.Println("The time is: ")
	now := time.Now()
	switch {

	case now.Hour() < 12:
		fmt.Println("Good morning")
	case now.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}

	fmt.Println("Counting asc")
	for i := 0; i < 10; i++ {
		fmt.Println("Counting ", i) // asc
	}
	fmt.Println("---------------------")
	fmt.Println("Counting desc")
	for i := 0; i < 10; i++ {
		defer fmt.Println("Counting ", i) // asc
	}

}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))

}

func pow(x, n, lim float64) float64 {
	if variable := math.Pow(x, n); variable > lim {
		return variable
	}
	return lim
}

//func Sqrt(x float64) float64 {
//	guess := 1.0
//	var variable float64 = guess
//
//	guess -= (guess*guess - x) / (2 * guess)
//	for guess < 10 {
//
//	}
//}
