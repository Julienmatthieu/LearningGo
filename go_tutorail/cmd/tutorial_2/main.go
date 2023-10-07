package main

import "fmt"

func main() {
	// Arrays : Fixed Length, Same Type, Indexable, Contiguous in Memory
	intArray := [...]int32{1, 2, 3}
	//intArray := [3]int32{1, 2, 3}
	//var intArray [3]int32 = [3]int32{1, 2, 3}
	//	var intArray [3]int32

	fmt.Println(intArray)

	/*	intArray[1] = 123
		fmt.Println(intArray[0])
		fmt.Println(intArray[1:3])

		fmt.Println(&intArray[0]) // & -> memory adresse
		fmt.Println(&intArray[1])
		fmt.Println(&intArray[2])
	*/

	// Slice is a Array but with more options
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with a capaticy of %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with a capaticy of %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// Maps {key, value}
	var myMap map[string]uint8 = make(map[string]uint8) // keys string value uint8
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Steven": 30}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"]) // <= Print la value par defaut de uint et pas d'erreur

	var age, ok = myMap2["Jason"] // ok boolean de présence
	if ok {
		fmt.Printf("The afe is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	// For loop
	// Pas d'ordre, Ce dernier peut etre différent entre chaque execution
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArray {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
