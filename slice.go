package main

import "fmt"

func main() {
	fmt.Println("Welcome to the slices")
	names := []string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
	fmt.Println(names[0:2]) // the startIndex include but the lestIndex Not
}
