package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type key string

//010_OMIT
func main() {
	http.HandleFunc("/dbl", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		r = r.WithContext(ctx)
		dbl(w, r) // HL
	})

	log.Println("server starting...")
	log.Fatal(http.ListenAndServe(":4998", nil))
}

//020_OMIT

//030_OMIT
func dbl(w http.ResponseWriter, r *http.Request) {
	result := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		i, err := strconv.Atoi(r.FormValue("q"))
		if err != nil {
			log.Printf("bad atoi: %v", r.FormValue("q"))
			return
		}
		result <- 2 * i
	}()
	select {
	case i := <-result:
		fmt.Fprintf(w, "%v\n", i) // HL
	case <-r.Context().Done():
		fmt.Fprintf(w, "timed out\n") // HL
	}
}

//040_OMIT
