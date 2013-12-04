package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

// declaring a map of string, Vertex pairs
var medical map[string]Vertex

// declaring a map of String, Vertex pairs with literal values
var computer = map[string]Vertex{
  "Bell Labs": Vertex{ 40.68433, -74.39967 },
  "Google": Vertex{ 37.42202, -122.08408 },
}

// declaring a map of String, Vertex pairs with literal values
// You can leave out the type in the Vertex literals as it is implied

/*
var computer = map[string]Vertex{
  "Bell Labs": {40.68433, -74.39967},
  "Google":    {37.42202, -122.08408},
}
*/

func main() {
  // allocate space to the map with make - note no size
  medical = make(map[string]Vertex)

  medical["Iora Health"] = Vertex{ 40.68433, -74.39967 }
  computer["Yahoo!"] = Vertex{ 40.68433, -74.39967 }

  delete(medical, "Iora Health")

  // 'comma ok' idiom
  v, ok := medical["Iora Health"]
  fmt.Println("The value:", v, "Present?", ok)

  // no := here (ok exists)
  v, ok = computer["Bell Labs"]
  fmt.Println("The value:", v, "Present?", ok)

  // assignment in if
  if v, ok = computer["Bell Labs"]; ok {
	  fmt.Println("The value: Bell Labs", v)
  } else {
	  fmt.Println("There is no value for: Bell Labs")
  }
}
