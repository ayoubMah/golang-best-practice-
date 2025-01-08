package main

import "fmt"

// Pointers

func incr(p *int) int {
	*p++ // increment the  value that p pointed on not the pointer itself
	return *p
}

func main() {
	//i, j := 42, 2701
	//
	//p := &i         // point to i
	//fmt.Println(*p) // read i through the pointer
	//*p = 21         // set i through the pointer
	//fmt.Println(i)  // see the new value of i
	//
	//p = &j         // point to j
	//*p = *p / 37   // divide j through the pointer => 2701/31 = 73
	//fmt.Println(j) // see the new value of j
	x := 1
	fmt.Println(incr(&x)) // will return 2
	fmt.Println(incr(&x)) // will return 3
}
