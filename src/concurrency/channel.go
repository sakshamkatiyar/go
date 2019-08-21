package main

import "fmt"

func main() {
	// Declare a unbuffered channel
	counter := make(chan int)
	// counter := make(chan int, 2) // no effect
	// Creates a buffered channel with capacity of 3
	nums := make(chan int, 3)
	// counter <- 2 // fatal error -> deadlock
	go func() {
		// Send value to the unbuffered channel
		counter <- 1
		counter <- 2
		close(counter) // Closes the channel
		// counter <- 5 // no effect
	}()

	go func() {
		// Send values to the buffered channel
		nums <- 10
		nums <- 30
		// close(nums)
		nums <- 40 // panic
		// nums <- 50 // deadlock
	}()
	// Read the value from unbuffered channel
	fmt.Println(<-counter)
	val, ok := <-counter // Trying to read from closed channel // blocking operation
	if ok {
		fmt.Println(val) // This won't execute
		// fmt.Println("val", val)
	}
	// Read the 3 buffered values from the buffered channel
	fmt.Println(<-nums)
	fmt.Println(<-nums)
	fmt.Println(<-nums)
	// fmt.Println(<-nums)
	close(nums) // Closes the channel
}