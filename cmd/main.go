package main

import (
        "github.com/chris530/gostuff/pkg/chris"
        "fmt"
        "sync"
        "time"
       )


type Duration int64

func mychan(t int, c chan string, s string, wg *sync.WaitGroup){ 

    defer wg.Done()
    duration := int(t)
    time.Sleep(time.Duration(duration) * time.Second) 
    fmt.Println(s, "from goroute")
    c <- s

} 


       
func main(){

        chris.Hello()

        var wg sync.WaitGroup
        c := make(chan string)

        wg.Add(3)
        go mychan(10, c, "chris", &wg)
        go mychan(30, c, "bill", &wg)
        go mychan(5, c, "nick", &wg)


        fmt.Println(<-c)
        fmt.Println(<-c)
        fmt.Println(<-c)

        wg.Wait()
   
}
