package main

import (
	"fmt"
	"time"
)

//010_OMIT
func main() {
	go disp("+", 3)
	go disp(".", 2)
	fmt.Println("main start")
	for i := 1; i <= 4; i++ { // try reducing to 3 or 2
		time.Sleep(time.Second)
	}
	time.Sleep(10 * time.Millisecond) // clean-up
	fmt.Println("\nmain end")
}

//020_OMIT
//030_OMIT
func disp(s string, n int) {
	for i := 1; i <= n; i++ {
		time.Sleep(time.Second)
		fmt.Print(s)
	}
}

//040_OMIT
