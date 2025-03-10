package main

import "fmt"

/*

Assignment
Complete the printReports function. It takes as input a sequence of messages, intro, body, outro. It should call printCostReport once for each message.

For each call of printCostReport, give it an anonymous function that returns the cost of a message as an integer. Here are the costs:

Intro: 2x the message length
Body: 3x the message length
Outro: 4x the message length
Use the built-in len() function to get the length of a string:

helloLen := len("hello")
// helloLen = 5
*/

func getProductMessage(tier string) string {
	quantityMsg, priceMsg, _ := getProductInfo(tier)
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

// don't touch below this line

func getProductInfo(tier string) (string, string, string) {
	if tier == "basic" {
		return "1,000 texts per month", "$30 per month", "most popular"
	} else if tier == "premium" {
		return "50,000 texts per month", "$60 per month", "best value"
	} else if tier == "enterprise" {
		return "unlimited texts per month", "$100 per month", "customizable"
	} else {
		return "", "", ""
	}

}

func getCoords() (x, y int) {
	// x and y are initialized with zero values

	return // automatically returns x and y
}

func printReports(intro, body, outro string) int {
	// Anonymous function for calculating costs
	printCostReport(func(msg string) int {
		return 2 * len(msg)
	}, intro)

	printCostReport(func(msg string) int {
		return 3 * len(msg)
	}, body)

	printCostReport(func(msg string) int {
		return 4 * len(msg)
	}, outro)

	return 0 // Return value is not specified in the prompt
}

func main() {
	fmt.Println(getCoords())
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
	fmt.Println(calculate(12, 12))
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

func calculate(a float64, b float64) float64 {
	return a * b
}
