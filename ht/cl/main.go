package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//010_OMIT
func main() {
	// setup -- create request
	const url = "http://localhost:4998/dbl"
	req, err := http.NewRequest("GET", url+"?q=2", nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	req = req.WithContext(ctx)
	defer cancel()

	// make the request
	log.Println("request started")
	httpDo(req)
	log.Println("request ended")
}

//020_OMIT

//030_OMIT
func httpDo(req *http.Request) {
	ch := make(chan string)
	cl := &http.Client{}
	go func() {
		res, err := cl.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		val, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		ch <- fmt.Sprintf("%s", val)
	}()
	select {
	case r := <-ch:
		fmt.Printf("recieved from server: %s", r)
	case <-req.Context().Done():
		fmt.Println("received Done(): request cancelled")
	}
}

//040_OMIT
