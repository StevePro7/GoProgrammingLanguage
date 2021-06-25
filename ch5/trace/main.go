package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOPeration() {
	defer trace("bigSlowOPeration")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit  %s (%s)", msg, time.Since(start))
	}
}

func main() {
	fmt.Println("Start")
	bigSlowOPeration()
	fmt.Println("-END-")
}
