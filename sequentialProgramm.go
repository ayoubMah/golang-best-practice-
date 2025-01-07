package main

import (
	"fmt"
	"time"
)

func calculateSquart(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
}

func main() {

	start := time.Now()
	for i := 1; i < 1000; i++ {
		go calculateSquart(i) //executed in a go routine, this will be executed in 1000 different go routines
	}
	elapsed := time.Since(start)
	time.Sleep(2 * time.Second)
	fmt.Printf("elapsed: %s\n", elapsed)

	//go start()
	//time.Sleep(1 * time.Second) // to make sure all go routines completed after run the Main
	go func() {
		fmt.Println("In anonymous function")
	}()
	time.Sleep(1 * time.Second)
}

// let's understand the Main GO Routine

func start() {
	go process()
	fmt.Println("in start")

}

func process() {
	fmt.Println("in process")

}
