package main

import "fmt"

func main() {
	text := "hello you"

	printMe(text)

	numerator := 11
	denominator := 2
	var result, remain int = intDivision(numerator, denominator)
	fmt.Printf("The result of the division is %v and the remain is %v", result, remain)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remain int = numerator % denominator

	return result, remain
}
