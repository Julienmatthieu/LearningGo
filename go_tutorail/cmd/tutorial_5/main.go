package main

import "fmt"

// Pointer: *
func main() {
	/*	var p *int32 = new(int32)
		var i int32

		fmt.Printf("\nThe value p point to is: %v", *p)
		fmt.Printf("\nThe value if i is : %v", i)

		p = &i
		*p = 10 // change the value at the position p
		fmt.Printf("\nThe value p point to is: %v", *p)
		fmt.Printf("\nThe value if i is : %v", i)*/

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nMemory location of thing1 is: %p", &thing1)
	var result [5]float64 = square(&thing1)

	fmt.Printf("\nThe resulr is: %v", result)
}

// ici nous avons une copie de l'array
// Avec un pointeur la duplication n'a pas lieu
func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nMemory location of thing2 is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
