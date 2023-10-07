//String

package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"             // Sera utilisé et itéré comme un array of bite
	var myRuneString = []rune("résumé") // rune equivalent a char

	fmt.Println(myString)
	var indexed = myString[0]
	fmt.Printf("%v %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe lenght of 'myRuneString' is %v", len(myRuneString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune) // ascii

	var strSlice = []string{"a", " ", "t", "e", "s", "t"}
	var strBuilder strings.Builder // Just like in  C#
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
