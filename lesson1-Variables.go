// in this lesson you will learn all what you need about variables in Golang
// enjoy !

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// stage1
	fmt.Println("Starting Textio server...")
	//Declaring a variable the sad way
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// stage2
	// GOATed variable declaration:
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"

	fmt.Println(messageStart, age, messageEnd)

	//stage3
	/*
		We are increasing the maximum message length from 140 to 280 characters.
		Very reluctantly, I might add.
		Users actually want to write more than 140 characters?!? Madness.
	*/
	maxMessageLength := 140
	newMaxMessageLength := 280
	fmt.Println("Textio is increasing the maximum message length from", maxMessageLength, "to", newMaxMessageLength, "characters.")

	//stage8
	accountAgeFloat := 2.6
	accountAgeInt := int(accountAgeFloat) // casting in golang

	fmt.Println("Your account has existed for", accountAgeInt, "years")

	// stage9
	var ayoub string = "fiin"
	var mah string = "cv"
	var aa string = ayoub + mah
	fmt.Println("ss: ", aa)

	//stage 10
	// same line declarations
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"

	fmt.Println(averageOpenRate, displayMessage)

	//satge 20

	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent \n", name, openRate)

	// %.1f    to replace float followed by 1 digit after the float
	// %s    to replace strings
	// %d    to replace integers

	fmt.Print(msg)

	//stage 21
	const nom = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(nom))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(nom))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", nom)

	//finale stage

	fname := "Dalinar"
	lname := "Kholin"
	agee := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	// Don't touch above this line

	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %v, Message: %s \n", fname, lname, agee, messageRate, isSubscribed, message)

	// Don't touch below this line

	fmt.Println(userLog)

}
