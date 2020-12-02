//
// Starting to notifing cool...
// Starting to notifing is...
// Finished notifying is...
// Starting to notifing super...
// Starting to notifing chris...
// Finished notifying cool...
// Finished notifying chris...
// Finished notifying super...
// All services notified!
//
//



package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func notify(services ...string) {
	var wg sync.WaitGroup

	for _, service := range services {
		wg.Add(1)
		go func(s string) {
			fmt.Printf("Starting to notifing %s...\n", s)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Printf("Finished notifying %s...\n", s)
			wg.Done()
		}(service)
	}

	wg.Wait()
	fmt.Println("All services notified!")
}


func main() {
	notify("chris" , "is", "super", "cool")
}
