package main

import (
        "fmt"
        "github.com/pkg/profile"
       )

func add_100(c chan int, n int){

    n = n + 100
    c <- n
}


func main(){

    defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

    cnt := 0
    c1 := make(chan int)

    for { 

      go add_100(c1,cnt)
      fmt.Println( <- c1 )
      cnt++
 
      if cnt > 100000000 {  
        break
      }

    }

}
