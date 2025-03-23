package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("Enter your name and age: :")
	count, err := fmt.Scanln("%s %d", &name, &age)
	fmt.Println("count : ", count)
	fmt.Println("err : ", err)
}
