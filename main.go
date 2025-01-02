package main

import (
	"fmt"
	"sync"
)

/*
Design and implement a Golang program that prints numbers from 1 to 10 in the correct sequence,
with even numbers printed by one goroutine and odd numbers printed by a separate goroutine.

We need to synchronize these go routines so that the numbers are printed in order.

Output: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10.
*/

func evenWorker(wg *sync.WaitGroup, oddChannel chan bool, evenChannel chan bool) {
	defer wg.Done()
	for i := 2; i <= 10; i = i + 2 {
		<-evenChannel
		fmt.Println(i)
		oddChannel <- true
	}
}

func oddWorker(wg *sync.WaitGroup, oddChannel chan bool, evenChannel chan bool) {
	defer wg.Done()
	for i := 1; i <= 10; i = i + 2 {
		<-oddChannel
		fmt.Println(i)
		evenChannel <- true
	}
}

func main() {
	fmt.Println("Printing numbers using go routines")
	var wg sync.WaitGroup
	oddChannel := make(chan bool, 1)
	evenChannel := make(chan bool, 1)
	defer close(evenChannel)
	defer close(oddChannel)
	wg.Add(2)
	go evenWorker(&wg, oddChannel, evenChannel)
	go oddWorker(&wg, oddChannel, evenChannel)
	oddChannel <- true
	wg.Wait()
	fmt.Println("Finished!!")
}
