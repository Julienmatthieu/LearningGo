package main

import (
	"errors"
	"fmt"
)

func main() {
	numerator := 12
	denominator := 3
	var result, remain, err = intDivision(numerator, denominator)
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remain != 0:
		fmt.Printf("The result of the division is %v and the remain is %v\n", result, remain)
	default:
		fmt.Printf("The result of the division is %v \n", result)
	}

	switch remain {
	case 0:
		fmt.Printf("The divison was exact")
	case 1, 2:
		fmt.Printf("The divison was close")
	default:
		fmt.Printf("The divison was not close")
	}
}

/*
func printMe(printValue string) {
	fmt.Println(printValue)
}*/

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remain int = numerator % denominator

	return result, remain, err
}
