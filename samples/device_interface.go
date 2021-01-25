package main

import (
	"fmt"
)

type DrawPower interface {
  Draw()
  PowerOn()
}


func(d device) Draw() {
  fmt.Println("Drawing power from wall socket")  
}

func(d device) PowerOn() {
  fmt.Println("PowerOn")  
}  
	
	
type device struct{
  name string
  devicetype string
  model string
} 



func main() {
	
  b := device{name: "Bender", devicetype: "blender",  model: "ADJDD" }    
  b.Draw()
  fmt.Println(b.name," initalized")

  f := device{name: "Fridge", devicetype: "fridge",  model: "HJFSXF" }    
  f.Draw()
  fmt.Println(f.name," initalized")

  w := device{name: "Washer", devicetype: "washer", model: "ZMXD" }    
  w.PowerOn()
  fmt.Println(w.name," initalized")
  

}
