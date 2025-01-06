// Conditionals

package main

import "fmt"

// stage 4: practise switch statement
func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0

	default:
		return 0.0
	}
}

func main() {
	//stage 2
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// don't touch above this line

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	//stage 4: practise switch statement
	plan := "basic"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "pro"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "enterprise"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "free"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "unknown"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))

	// Exercice:

	/*

		Using the given variables, write conditional statements to calculate and update the variables.

		First, set finalCost to the bulkMessageCost. If the user is premium, then the discountRate should be applied to the finalCost. For example, a discountRate of 0.10 means the discounted price per message would be 90% of the original price.

		Next, if the user has enough money in their accountBalance:

		Process the payment by deducting the finalCost from their accountBalance
		Print the purchaseSuccessMessage
		Otherwise, just print the insufficientFundMessage.

	*/

	// please remove the comment  from these variables

	//var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	//var purchaseSuccessMessage string = "Purchase successful."
	//var accountBalance float64 = 100.0
	//var bulkMessageCost float64 = 75.0
	//var isPremiumUser bool = true
	//var discountRate float64 = 0.10
	//var finalCost float64

	// write your code here

	//fmt.Println("Account balance:", accountBalance)

}
