package main

import (
	"fmt"
	"time"
)

// isIAmdone function runs in a Goroutine and checks every 30 seconds
// if the current minute is odd. If it is odd, it sends true on the done channel.
func isIAmdone(done chan bool) {
	for {
		time.Sleep(30 * time.Second)
		if time.Now().Minute()%2 != 0 {
			done <- true
		}
	}
}

func main() {
	done := make(chan bool)
	// Start the isIAmdone function in a Goroutine
	go isIAmdone(done)

	// Infinite loop in the main function
	for {
		select {
		case <-done:
			fmt.Println("isIAmdone returned true. Terminating main loop.")
			return
		default:
			// Do some work in the main loop
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}
