// In this go code you gana learn GO HHHHH
// just following me
// or better go to     https://go.dev/tour
// happy codding :)

package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("The time is", time.Now())           // time.Now() return the current time
	fmt.Println("A random number is", rand.Intn(10)) // return a random number between O and 10 (10 not included)

	// so what i understand is that the %g replaced with the operation a condition of Printf not Println
	fmt.Printf("The sqrt of 16 is : %g right \n", math.Sqrt(16))

	// this line should be return an error right cuz pi starts with a small capital which is unexported name
	//fmt.Println("the pi equals to", math.pi)

	// to fix the error we should replace the pi with Pi (a capital lettre)
	fmt.Println("the pi equals to", math.Pi) // so this will give 3.14...

	// so let's call our add function to calculate the some
	a := add(23, 56)
	fmt.Println(" 23 + 56 = ", a)
	// or you can simply do this: fmt.Println(" 23 + 56 = ", add(23,56))

	b := add3Numbers(10, 20, 30)
	fmt.Println("10+20+30 = ", b)
	// the same thing here

	c, d := swap(12, "Hello World")
	fmt.Println(c, d)

	// use naked return
	e, r := split(16)
	fmt.Println("The split function return : ", e, " and ", r)

	// or just like this
	fmt.Println(split(12))

	var java, python, rust bool
	var i int
	var k, l int = 23, 33           // declare and initialize variables
	fmt.Println(java, python, rust) // return false false false
	fmt.Println(i)                  // return 0
	fmt.Println(k, l)

	// declare and initialize variables
	qo, do, mo := 3, true, "GI"
	fmt.Println(qo, do, mo)

	var (
		ToBe   bool       = true
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Println(ToBe, MaxInt, z)
	fmt.Printf("Type:  %T , value of , %v \n", ToBe, ToBe)
	fmt.Printf("Type:  %T , value of , %v \n", MaxInt, MaxInt)
	fmt.Printf("Type:  %T , value of , %v \n", z, z)

	// Zero  values  or default values,
	var num int
	var bol bool
	var fl float64
	var st string
	fmt.Printf("%v %v %v %v \n", num, fl, st, bol) // why %q and not %v ??
	// so yeah i understand why %q and not %v
	// so %v give you the value and the %q it outputs the string enclosed in double quotes

	// type Inferance
	ayoub := 12 + 34i
	fmt.Printf("The type is: %T of the value: %v \n", ayoub, ayoub)
	// just change the value of ayoub and you'll get different type

	//Constants:
	// Constants cannot be declared using the := syntax . this is very important !!!!!!!!

	const Pi = math.Pi
	const World = "World"
	fmt.Println(Pi)
	fmt.Println(World)

	// Numeric Constants
	const (
		Big   = 1 << 99
		Small = Big >> 99
	)
	//fmt.Println(Big)
	fmt.Println(calo(Small), "ssss", cala(Big))

}

func calo(x int) int         { return x + 1 }
func cala(y float64) float64 { return y * 0.1 }

// so let's define a function that calculates the some of 2 numbers
// remember in GO the type came after the  variables and the functions
func add(x int, y int) int {
	return x + y
}

// so we can do x int, y int or we can do also : x,y int
func add3Numbers(x, y, z int) int {
	return x + y + z

}

func swap(x int, y string) (int, string) {
	return x, y
}

// naked return
func split(sum int) (w, z int) {
	w = sum + 9/9
	z = sum % 2
	return
}
