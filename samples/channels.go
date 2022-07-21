// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func mytest(yo *int, c chan int) {

	fmt.Println("I am in the function ", *yo)
	*yo = *yo + 7

	// send back 14
	c <- *yo

}

func main() {
	c := make(chan int, 1)
	var p int = 7
	go mytest(&p, c)
	n := <-c
	fmt.Println("Back in main ", n)

	// pointer set p to 17 also
	fmt.Println("new p", p)
	fmt.Scanln()
}
