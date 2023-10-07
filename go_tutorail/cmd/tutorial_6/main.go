package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{} // short for Mutual Exclusion
var waitGroup = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var resulr = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		waitGroup.Add(1) // Ajoute une compteur au waitgroupe
		go dbCall(i)     //Go key word to use different task
	}
	waitGroup.Wait() // attend que toutes les tasks soient terménies
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", resulr)
}

func dbCall(i int) {
	// Simulate DB cal delays
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	m.Lock()
	resulr = append(resulr, dbData[i])
	m.Unlock()
	waitGroup.Done() // indique ai waitgroup que la task est terminée
}
