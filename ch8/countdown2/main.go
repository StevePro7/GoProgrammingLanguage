package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Lift off!")
}

func main() {

	// create abort channel
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a signle byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")

	select {
	case <-time.After(10 * time.Second):
		// Do nothing
	case <-abort:
		fmt.Println("Launch aborted")
		return
	}

	launch()
}
