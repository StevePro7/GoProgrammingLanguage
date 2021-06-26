package main

import (
	"fmt"
	"log"
)

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			log.Printf("C x = %d  ", x)
			naturals <- x
			log.Printf("C N = %d  ", x)
			log.Println()
		}
	}()

	// Squares
	go func() {
		for {
			log.Printf("S N1= %d  ", naturals)
			log.Printf("S S1= %d  ", squares)
			x := <-naturals
			squares <- x * x
			log.Printf("S N2= %d  ", naturals)
			log.Printf("S S2= %d  ", squares)
		}
	}()

	// Printer (in main goroutines)
	log.Print("MAIN squares")
	fmt.Println(<-squares)
}
