package main

import (
	"fmt"
	"sync"
	"time"
)


func a(wg *sync.WaitGroup){

  time.Sleep(5 * time.Second)  // sleep 5 sec
  fmt.Println("Running a")
  wg.Done()

}


func b(wg *sync.WaitGroup){

  fmt.Println("Running b")
  wg.Done()

 
}


func main() {

  
  var wg sync.WaitGroup
  wg.Add(2)

  go a(&wg)
  go b(&wg)

  wg.Wait()

  fmt.Println("I am done ")
	
}
