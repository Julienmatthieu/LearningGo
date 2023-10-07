package main

import "fmt"

// Channels : Hold Date, Thread safe, listen for data

func main() {
	var c = make(chan int)
	go process(c)
	for i := range c {
		fmt.Println(i)
		// Ici l'on va print a chaque fois qu'une value est ajouté au channel
		// Dans notre cas on ajoute dans la fonction process
	}
}

func process(c chan int) {
	defer close(c) // defer => do something juste before the fonction exit
	for i := 0; i < 5; i++ {
		c <- i
	}
	//close(c) // close the channel
}
