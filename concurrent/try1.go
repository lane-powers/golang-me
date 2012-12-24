package main

// playing with the base concurrency idea from the golang tutorial

import "fmt"

func fibonacci(c, quit chan int, n int) {
	x, y := 0, 1
	t := 0
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			t++
			if t >= n {
				// all goRoutines are finished
				fmt.Println("quit")
				return
			}
		}
	}
}

func main() {
	nuRots := 5
	c := make(chan int)
	quit := make(chan int)
	// spin us up nuRots counters
	for l := 0; l < nuRots; l++ {
		go func() {
			fmt.Println("In go routine:")
			for i := 0; i < 10; i++ {
				fmt.Println(<-c)
			}
			quit <- 0
		}()
	}
	fibonacci(c, quit, nuRots)
}