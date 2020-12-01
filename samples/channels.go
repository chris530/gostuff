//
// go run mytest.go
// Two
// Three
// One



package main

import (
        "fmt"
        "time"
       )


func run1(cc chan string){
  time.Sleep(7 * time.Second)
  cc<-"One"
}

func run2(cc chan string){
  time.Sleep(1 * time.Second)
  cc<-"Two"
}

func run3(cc chan string){
  time.Sleep(4 * time.Second)
  cc<-"Three"
}


func main(){

  c := make(chan string)
  
  go run1(c)
  go run2(c)
  go run3(c)

  fmt.Println(<-c) 
  fmt.Println(<-c) 
  fmt.Println(<-c) 

}
