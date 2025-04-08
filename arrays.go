package main

import "fmt"

func main() {
	var args [5]int = [5]int{1, 2, 3, 4}

	fruits := [...]string{"Apple", "Orange", "Banana"}

	matrex := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	args[0] = 19

	fmt.Println(fruits)
	fmt.Println(args)
	fmt.Println(len(fruits))
	fmt.Println(len(args))
	fmt.Println(args[0])
	//fmt.Println(fruits[2])
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	for index, elm := range args {
		fmt.Println(index, " => ", elm)
	}

	fmt.Println("getting the 9 value ")
	fmt.Println(matrex[2][2])

}
