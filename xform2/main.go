package main

import (
	"fmt"
	"time"
)

//010 OMIT
func main() {
	done := make(chan bool)
	n := 3
	c1 := gen(n) // generate n integers
	c2 := dbl("a", done, c1)
	c3 := dbl("b", done, c2)
	for i := 1; i <= n; i++ {
		fmt.Println(<-c3)
	}
	done <- true // ; done <- true
	//close(done)
	// Do other work.
	time.Sleep(10 * time.Millisecond)
}

//020 OMIT
//030 OMIT
func gen(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

//040 OMIT
//050 OMIT
func dbl(name string, done <-chan bool, i <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for { // note the forever loop
			select {
			case v := <-i:
				ch <- 2 * v
			case s := <-done: // HL
				fmt.Println(name, "shutting down, received:", s) // HL
				return
			}
		}
	}()
	return ch
}

//060 OMIT
