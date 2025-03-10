package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Converting String to int using Atoi()")
	var s string = "123a" // correct the error by removing 'a'   "123a" ---> "123"

	i, err := strconv.Atoi(s)
	fmt.Printf("%v , %T \n", i, i)
	fmt.Printf("%v , %T \n", err, err)

	fmt.Println("====================================================")

	fmt.Println("Converting int to string using Itoa() ")
	var a int = 244
	st := strconv.Itoa(a)
	fmt.Printf("%v , %T \n", st, st)

}
