package main

// Chapter 2: Variables
//func main() {
//	// 2.3. Type Inference
//	//penniesPerText := 2.0
//	//fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
//
//	// 2.4. Same Line Declaration
//	//averageOpenRate, displayMessage := .23, "is the average open rate of your message"
//	//fmt.Println(averageOpenRate, displayMessage)
//
//	// 2.5. Type Sizes
//	//accountAge := 2.6
//	//accountAgeInt := int(accountAge)
//	//fmt.Println("Your account has existed for", accountAgeInt, "years")
//
//	// 2.8. Constants
//	//const premiumPlanName = "Premium Plan"
//	//const basicPlanName = "Basic Plan"
//	//fmt.Println("plan:", premiumPlanName)
//	//fmt.Println("plan:", basicPlanName)
//
//	// 2.9. Computed Constants
//	//const secondsInMinute = 60
//	//const minutesInHour = 60
//	//const secondsInHour = secondsInMinute * minutesInHour
//	//fmt.Println("number of seconds in an hour:", secondsInHour)
//
//	// 2.10. Formatted String In Go
//	//const name = "Saul Goodman"
//	//const openRate = 30.5
//	//// Wrong solution - string const cannot have float value
//	////const msg = "Hi " + name + ", your open rate is " + string(openRate) + " percent"
//	//
//	//// Correct solution
//	//msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)
//	//fmt.Println(msg)
//
//	// 2.11. Conditional
//	//messageLen := 10
//	//maxMessageLen := 20
//	//fmt.Println("Trying to send a message if length:", messageLen)
//	//if messageLen <= maxMessageLen {
//	//	fmt.Println("Message sent")
//	//} else {
//	//	fmt.Println("Message not sent")
//	//}
//}

// Chapter 3: Functions
// 3.1. Functions
//func concat(s1 string, s2 string) string {
//	return s1 + s2
//}
//func main() {
//	fmt.Println(concat("Hello ", "World!"))
//	fmt.Println(concat("My favourite Go book is ", "\"Let's Go\" by Alex Edwards"))
//}

// 3.6. Passing Variables By Values
//func main() {
//	sendsSoFar := 430
//	const sendsToAdd = 25
//	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
//	fmt.Println("You've sent ", sendsSoFar, " messages")
//}
//
//func incrementSends(sendsSoFar, sendsToAdd int) int {
//	sendsSoFar += sendsToAdd
//	return sendsSoFar
//}

// 3.7. Ignoring Return Values
//func main() {
//	firstName, _ := getName()
//	fmt.Println("Welcome to Springfield,", firstName+"!")
//}
//
//func getName() (string, string) {
//	return "Lisa", "Simpson"
//}

// 3.8. Named Return Value
// In a function with no explicit return argument will result
// in that function returning named return value, which is
// declared in function signature
//func yearsUntilEvents(age int) (
//	yearsUntilAdult int,
//	yearsUntilDrinking int,
//	yearsUntilCarRental int,
//) {
//	yearsUntilAdult = 18 - age
//	if yearsUntilAdult < 0 {
//		yearsUntilAdult = 0
//	}
//	yearsUntilDrinking = 21 - age
//	if yearsUntilDrinking < 0 {
//		yearsUntilDrinking = 0
//	}
//	yearsUntilCarRental = 25 - age
//	if yearsUntilCarRental < 0 {
//		yearsUntilCarRental = 0
//	}
//	return
//}
//
//func main() {
//	println(yearsUntilEvents(22))
//}
