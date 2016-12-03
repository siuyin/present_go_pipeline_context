package main

import (
	"context"
	"log"
	"os/exec"
	"time"
)

//010_OMIT
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	log.Println("started...")
	go func() {
		time.Sleep(1 * time.Second) // Try 3 secs // HL
		cancel()
	}()
	cmd := exec.CommandContext(ctx, "sleep", "2")
	if err := cmd.Run(); err != nil {
		log.Printf("command run: %v", err)
	}
	log.Println("ended.")
}

//020_OMIT
