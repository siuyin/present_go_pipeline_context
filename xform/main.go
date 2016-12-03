package main

import (
	"fmt"
	"time"
)

//010 OMIT
func main() {
	n := 3
	c1 := gen(n) // generate n integers
	c2 := dbl(c1)
	//c3 := dbl(c2) // HL
	for i := 1; i <= n; i++ {
		fmt.Println(<-c2) // uncomment c3 and read from c3 // HL
	}
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
func dbl(i <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for { // note the forever loop
			ch <- 2 * <-i
		}
	}()
	return ch
}

//060 OMIT
